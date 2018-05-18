package cluster

import (
	"fmt"
	"io/ioutil"
	stdlog "log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/hashicorp/serf/serf"
	"github.com/rs/zerolog/log"
	"github.com/ubiqueworks/joat/util"
)

const (
	RoleController = "controller"
	RoleWorker     = "worker"
)

func NewClusterManager(conf *Config) Manager {
	manager := &serfClusterManager{
		advertiseIP:   conf.AdvertiseIP,
		clusterAddr:   conf.ClusterAddr,
		dataDir:       filepath.Join(conf.DataDir, ".cluster"),
		enableSerfLog: conf.EnableSerfLog,
		eventCh:       make(chan serf.Event, 256),
		members:       make(map[string]*MemberInfo),
		role:          conf.Role,
		rpcPort:       conf.RpcPort,
		shutdownCh:    make(chan struct{}),
	}
	return manager
}

type MemberInfo struct {
	ID       string `json:"id"`
	Role     string `json:"role"`
	Addr     string `json:"address"`
	SerfPort int    `json:"serf_port"`
	RpcPort  int    `json:"rpc_port"`
	Status   string `json:"status"`
}

func (m *MemberInfo) ClusterAddr() *net.TCPAddr {
	return &net.TCPAddr{IP: net.ParseIP(m.Addr), Port: m.SerfPort}
}

func (m *MemberInfo) RpcAddr() *net.TCPAddr {
	return &net.TCPAddr{IP: net.ParseIP(m.Addr), Port: m.RpcPort}
}

type Config struct {
	AdvertiseIP   *net.IP
	ClusterAddr   *net.TCPAddr
	DataDir       string
	EnableSerfLog bool
	Role          string
	RpcPort       int
}

type Manager interface {
	Join(addr *net.TCPAddr) error
	Members(role string) []*MemberInfo
	Start() error
	Stop()
}

type serfClusterManager struct {
	advertiseIP   *net.IP
	clusterAddr   *net.TCPAddr
	dataDir       string
	enableSerfLog bool
	eventCh       chan serf.Event
	memberID      string
	members       map[string]*MemberInfo
	membersLock   sync.RWMutex
	role          string
	rpcPort       int
	serf          *serf.Serf
	shutdownLock  sync.Mutex
	shutdownCh    chan struct{}
	shutdown      bool
}

func (m *serfClusterManager) Members(role string) []*MemberInfo {
	m.membersLock.RLock()
	defer m.membersLock.RUnlock()

	members := make([]*MemberInfo, 0)
	for _, member := range m.members {
		if role == "" || member.Role == role {
			members = append(members, member)
		}
	}
	return members
}

func (m *serfClusterManager) Join(addr *net.TCPAddr) error {
	if _, err := m.serf.Join([]string{addr.String()}, true); err != nil {
		return err
	}
	return nil
}

func (m *serfClusterManager) Start() error {
	if err := util.EnsurePath(m.dataDir, true); err != nil {
		log.Error().Err(err).Msgf("error creating cluster data dir")
		return fmt.Errorf("error creating cluster data dir: %s", m.dataDir)
	}

	if err := m.ensureMemberID(); err != nil {
		log.Error().Err(err).Msg("error ensuring member id")
		return fmt.Errorf("error ensuring member id")
	}

	snapshotPath := filepath.Join(m.dataDir, "serf-snapshot")

	if err := util.EnsurePath(snapshotPath, false); err != nil {
		log.Error().Err(err).Msgf("error creating serf snapshot dir")
		return fmt.Errorf("error creating serf snapshot dir: %s", snapshotPath)
	}

	bindIP := m.clusterAddr.IP.String()
	serfAddr := &net.TCPAddr{
		IP:   *m.advertiseIP,
		Port: m.clusterAddr.Port,
	}

	serfConf := serf.DefaultConfig()
	serfConf.Init()
	serfConf.Tags["role"] = m.role
	serfConf.Tags["serf_port"] = fmt.Sprintf("%d", serfAddr.Port)
	serfConf.Tags["rpc_port"] = fmt.Sprintf("%d", m.rpcPort)

	serfConf.EventCh = m.eventCh
	serfConf.NodeName = m.memberID
	serfConf.SnapshotPath = snapshotPath
	serfConf.MemberlistConfig.BindAddr = bindIP
	serfConf.MemberlistConfig.BindPort = serfAddr.Port
	serfConf.MemberlistConfig.AdvertiseAddr = serfAddr.IP.String()
	serfConf.MemberlistConfig.AdvertisePort = serfAddr.Port
	serfConf.MemberlistConfig.LogOutput = ioutil.Discard

	if m.enableSerfLog {
		serfConf.Logger = stdlog.New(log.Logger, "", 0)
	} else {
		serfConf.LogOutput = ioutil.Discard
	}

	if s, err := serf.Create(serfConf); err != nil {
		log.Error().Err(err).Msgf("error initializing serf")
		return fmt.Errorf("error initializing serf")
	} else {
		m.serf = s
	}

	go m.monitorClusterEvents()

	return nil
}

func (m *serfClusterManager) Stop() {
	if m.serf != nil {
		m.serf.Leave()
		m.serf.Shutdown()
	}
	m.shutdownLock.Lock()
	defer m.shutdownLock.Unlock()

	if m.shutdown {
		return
	}
	m.shutdown = true
	close(m.shutdownCh)
}

func (m *serfClusterManager) ensureMemberID() error {
	if m.memberID != "" {
		return nil
	}

	memberIDFilePath := filepath.Join(m.dataDir, ".id")

	if _, err := os.Stat(memberIDFilePath); err == nil {
		data, err := ioutil.ReadFile(memberIDFilePath)
		if err != nil {
			return err
		}
		m.memberID = string(data)
	} else {
		memberID := fmt.Sprintf("%s-%s", m.role, util.RandomHexString(8))
		if err := ioutil.WriteFile(memberIDFilePath, []byte(memberID), 0644); err != nil {
			return err
		}
		m.memberID = memberID
	}
	return nil
}

func (m *serfClusterManager) monitorClusterEvents() {
	for {
		select {
		case e := <-m.eventCh:
			switch e.EventType() {
			case serf.EventMemberJoin:
				m.addMember(e.(serf.MemberEvent))
			case serf.EventMemberLeave, serf.EventMemberFailed:
				m.removeMember(e.(serf.MemberEvent))
			case serf.EventMemberUpdate:
				m.updateMember(e.(serf.MemberEvent))
			//case serf.EventMemberReap:
			//	//p.localMemberEvent(e.(serf.MemberEvent))
			//case serf.EventMemberUpdate, serf.EventUser, serf.EventQuery:
			//	// Ignore
			default:
				log.Warn().Msgf("unhandled serf event: %#v", e)
			}
		case <-m.shutdownCh:
			return
		}
	}
}

func (m *serfClusterManager) addMember(ev serf.MemberEvent) {
	for _, member := range ev.Members {
		ok, info := m.createMemberInfo(member)
		if !ok {
			log.Warn().Msg("Attempt to add a member with an unknown role to the pool")
			continue
		}

		m.membersLock.Lock()
		if _, exists := m.members[info.ID]; !exists {
			log.Info().Msgf("Adding member to pool: %s [%s]", info.ID, info.Addr)
			m.members[info.ID] = info
		}
		m.membersLock.Unlock()
	}
}

func (m *serfClusterManager) removeMember(ev serf.MemberEvent) {
	for _, member := range ev.Members {
		ok, info := m.createMemberInfo(member)
		if !ok {
			log.Warn().Msg("Attempt to remove a member with an unknown role to the pool")
			continue
		}

		m.membersLock.Lock()
		if _, exists := m.members[info.ID]; exists {
			log.Info().Msgf("Removing member from pool: %s [%s]", info.ID, info.Addr)
			delete(m.members, info.ID)
		}
		m.membersLock.Unlock()
	}
}

func (m *serfClusterManager) updateMember(ev serf.MemberEvent) {
	for _, member := range ev.Members {
		ok, info := m.createMemberInfo(member)
		if !ok {
			log.Warn().Msg("Attempt to update a member with an unknown role to the pool")
			continue
		}

		m.membersLock.Lock()
		if _, exists := m.members[info.ID]; exists {
			log.Info().Msgf("Updating member in pool: %s [%s]", info.ID, info.Addr)
			m.members[info.ID] = info
		}
		m.membersLock.Unlock()
	}
}

func (m *serfClusterManager) createMemberInfo(member serf.Member) (bool, *MemberInfo) {
	role := member.Tags["role"]
	if role != RoleController && role != RoleWorker {
		return false, nil
	}

	rpcPort, err := strconv.Atoi(member.Tags["rpc_port"])
	if err != nil {
		return false, nil
	}

	serfPort, err := strconv.Atoi(member.Tags["serf_port"])
	if err != nil {
		return false, nil
	}

	parts := &MemberInfo{
		ID:       member.Name,
		Addr:     member.Addr.String(),
		Role:     role,
		RpcPort:  rpcPort,
		SerfPort: serfPort,
		Status:   member.Status.String(),
	}
	return true, parts
}
