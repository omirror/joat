package cmd

import (
	"fmt"
	"net"

	"github.com/hashicorp/go-sockaddr/template"
	"github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ubiqueworks/joat/util"
	"github.com/ubiqueworks/joat/worker"
)

func init() {
	workerCmd.Flags().String(flagAdvertiseAddr, "", flagAdvertiseAddrDesc)
	workerCmd.Flags().String(flagBindAddr, defaultBindAddr, flagBindAddrDesc)
	workerCmd.Flags().Int(flagClusterPort, defaultClusterPort, flagClusterPortDesc)
	workerCmd.Flags().String(flagDataDir, defaultDataDir, flagDataDirDesc)
	workerCmd.Flags().String(flagJoin, "", flagJoinDesc)
	workerCmd.Flags().Int(flagRpcPort, defaultRpcPort, flagRpcPortDesc)
	workerCmd.MarkFlagRequired(flagJoin)
	rootCmd.AddCommand(workerCmd)
}

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start a worker instance",
	RunE:  workerRun,
}

func workerRun(_ *cobra.Command, _ []string) error {
	log.Info().Msg("Starting controller instance...")

	advertiseAddr := viper.GetString(flagAdvertiseAddr)
	if advertiseAddr == "" {
		resolvedAdvertiseAddr, err := template.Parse("{{ GetPrivateIP }}")
		if err != nil {
			return fmt.Errorf("failed to get advertise IP address: %v", err)
		}
		advertiseAddr = resolvedAdvertiseAddr
	}
	bindAddr := viper.GetString(flagBindAddr)

	advertiseIP := net.ParseIP(advertiseAddr)
	bindIP := net.ParseIP(bindAddr)

	clusterPort := viper.GetInt(flagClusterPort)
	if !util.IsValidPort(clusterPort) {
		return fmt.Errorf("invalid cluster port: %d", clusterPort)
	}

	rpcPort := viper.GetInt(flagRpcPort)
	if !util.IsValidPort(rpcPort) {
		return fmt.Errorf("invalid rpc port: %d", rpcPort)
	}

	joinAddr := viper.GetString(flagJoin)
	joinTCPAddr, err := net.ResolveTCPAddr("tcp", joinAddr)
	if err != nil {
		return fmt.Errorf("error resolving join address: %s", joinAddr)
	}

	dataDir, _ := homedir.Expand(viper.GetString(flagDataDir))

	config := &worker.Config{
		AdvertiseIP: &advertiseIP,
		ClusterAddr: &net.TCPAddr{IP: bindIP, Port: clusterPort},
		DataDir:     dataDir,
		JoinAddr:    joinTCPAddr,
		RpcAddr:     &net.TCPAddr{IP: bindIP, Port: rpcPort},
		VerboseMode: viper.GetBool(flagVerbose),
	}

	log.Debug().Msgf("advertise IP: %v", advertiseIP.String())
	log.Debug().Msgf("cluster addr: %v", config.ClusterAddr.String())
	log.Debug().Msgf("data dir: %v", config.DataDir)
	log.Debug().Msgf("join addr: %v", config.JoinAddr.String())
	log.Debug().Msgf("rpc addr: %v", config.RpcAddr.String())

	return worker.StartWorker(config)
}
