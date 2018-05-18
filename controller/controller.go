package controller

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/ubiqueworks/joat/cluster"
	"github.com/ubiqueworks/joat/datastore"
	"github.com/ubiqueworks/joat/rpc"
	"google.golang.org/grpc"
)

type Config struct {
	AdvertiseIP *net.IP
	ClusterAddr *net.TCPAddr
	DataDir     string
	VerboseMode bool
	HttpAddr    *net.TCPAddr
	RpcAddr     *net.TCPAddr
}

func StartController(conf *Config) error {
	controller := &controller{
		config:     conf,
		shutdownCh: make(chan struct{}, 1),
	}
	return controller.start()
}

type controller struct {
	config         *Config
	clusterManager cluster.Manager
	repo           *repository
	store          datastore.Store
	shutdownLock   sync.Mutex
	shutdownCh     chan struct{}
	shutdown       bool
}

func (c *controller) start() error {
	if dbStore, err := datastore.NewStore(&datastore.Config{DataDir: c.config.DataDir}); err != nil {
		return err
	} else {
		c.store = dbStore
	}

	errCh := make(chan error)
	shutdownCh := c.shutdownCh

	var wg sync.WaitGroup
	wg.Add(3)

	go c.startHttpServer(&wg, shutdownCh, errCh)
	go c.startRpcServer(&wg, shutdownCh, errCh)
	go c.startClusterManager(&wg, shutdownCh, errCh)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errCh:
		log.Error().Err(err).Msg("caught service error")
		c.stop()
	case <-quit:
		c.stop()
	}

	log.Debug().Msg("waiting for shutdown to complete...")
	wg.Wait()
	log.Debug().Msg("shutdown completed")

	return nil
}

func (c *controller) stop() {
	c.shutdownLock.Lock()
	defer c.shutdownLock.Unlock()

	if c.shutdown {
		return
	}
	c.shutdown = true
	close(c.shutdownCh)
}

func (c *controller) startHttpServer(wg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	defer wg.Done()

	startErrCh := make(chan error)
	doneCh := make(chan struct{}, 1)

	serverAddr := c.config.HttpAddr.String()
	server := &http.Server{
		Addr:    serverAddr,
		Handler: configureRouter(c),
	}

	startFunc := func() {
		log.Info().Msgf("Starting HTTP server on %v", serverAddr)
		if err := server.ListenAndServe(); err != nil {
			startErrCh <- err
		}
	}

	stopFunc := func() {
		log.Debug().Msg("Stopping HTTP server...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		server.Shutdown(ctx)
		close(doneCh)
	}

	go startFunc()

	select {
	case err := <-startErrCh:
		errCh <- err
		close(doneCh)
	case <-shutdownCh:
		stopFunc()
	}

	<-doneCh
	log.Info().Msg("HTTP server stopped")
}

func (c *controller) startRpcServer(wg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	defer wg.Done()

	startErrCh := make(chan error)
	doneCh := make(chan struct{}, 1)

	grpcServer := grpc.NewServer()
	rpc.RegisterControllerServiceServer(grpcServer, newRpcServer(c))

	startFunc := func() {
		serverAddr := c.config.RpcAddr.String()
		log.Info().Msgf("Starting RPC server on %v", serverAddr)
		listener, err := net.Listen("tcp", serverAddr)
		if err != nil {
			startErrCh <- err
			return
		}
		grpcServer.Serve(listener)
	}

	stopFunc := func() {
		log.Debug().Msg("Stopping RPC server...")
		grpcServer.GracefulStop()
		close(doneCh)
	}

	go startFunc()

	select {
	case err := <-startErrCh:
		errCh <- err
		close(doneCh)
	case <-shutdownCh:
		stopFunc()
	}

	<-doneCh
	log.Info().Msg("RPC server stopped")
}

func (c *controller) startClusterManager(wg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	defer wg.Done()

	startErrCh := make(chan error)
	doneCh := make(chan struct{}, 1)

	manager := cluster.NewClusterManager(&cluster.Config{
		AdvertiseIP:   c.config.AdvertiseIP,
		ClusterAddr:   c.config.ClusterAddr,
		Role:          cluster.RoleController,
		DataDir:       c.config.DataDir,
		EnableSerfLog: c.config.VerboseMode,
		RpcPort:       c.config.RpcAddr.Port,
	})
	c.clusterManager = manager

	startFunc := func() {
		serverAddr := c.config.ClusterAddr.String()
		log.Info().Msgf("Starting cluster manager on %v", serverAddr)
		if err := manager.Start(); err != nil {
			startErrCh <- err
		}
	}

	stopFunc := func() {
		log.Debug().Msg("Stopping cluster manager...")
		manager.Stop()
		close(doneCh)
	}

	go startFunc()

	select {
	case err := <-startErrCh:
		errCh <- err
		stopFunc()
	case <-shutdownCh:
		stopFunc()
	}

	<-doneCh
	log.Info().Msg("Cluster manager stopped")
}

func (c *controller) members(role string) []*cluster.MemberInfo {
	return c.clusterManager.Members(role)
}
