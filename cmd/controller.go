package cmd

import (
	"fmt"
	"net"

	"github.com/hashicorp/go-sockaddr/template"
	"github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ubiqueworks/joat/controller"
	"github.com/ubiqueworks/joat/util"
)

func init() {
	controllerCmd.Flags().String(flagAdvertiseAddr, "", flagAdvertiseAddrDesc)
	controllerCmd.Flags().String(flagBindAddr, defaultBindAddr, flagBindAddrDesc)
	controllerCmd.Flags().Int(flagClusterPort, defaultClusterPort, flagClusterPortDesc)
	controllerCmd.Flags().String(flagDataDir, defaultDataDir, flagDataDirDesc)
	controllerCmd.Flags().String(flagHttpBindAddr, defaultBindAddr, flagHttpBindAddrDesc)
	controllerCmd.Flags().Int(flagHttpPort, defaultHttpPort, flagHttpPortDesc)
	controllerCmd.Flags().Int(flagRpcPort, defaultRpcPort, flagRpcPortDesc)
	rootCmd.AddCommand(controllerCmd)
}

var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "Start a controller instance",
	RunE:  controllerRun,
}

func controllerRun(_ *cobra.Command, _ []string) error {
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
	httpBindAddr := viper.GetString(flagHttpBindAddr)

	advertiseIP := net.ParseIP(advertiseAddr)
	bindIP := net.ParseIP(bindAddr)
	httpServerIP := net.ParseIP(httpBindAddr)

	clusterPort := viper.GetInt(flagClusterPort)
	if !util.IsValidPort(clusterPort) {
		return fmt.Errorf("invalid cluster port: %d", clusterPort)
	}

	httpPort := viper.GetInt(flagHttpPort)
	if !util.IsValidPort(httpPort) {
		return fmt.Errorf("invalid http port: %d", httpPort)
	}

	rpcPort := viper.GetInt(flagRpcPort)
	if !util.IsValidPort(rpcPort) {
		return fmt.Errorf("invalid rpc port: %d", rpcPort)
	}

	dataDir, _ := homedir.Expand(viper.GetString(flagDataDir))

	config := &controller.Config{
		AdvertiseIP: &advertiseIP,
		ClusterAddr: &net.TCPAddr{IP: bindIP, Port: clusterPort},
		DataDir:     dataDir,
		HttpAddr:    &net.TCPAddr{IP: httpServerIP, Port: httpPort},
		RpcAddr:     &net.TCPAddr{IP: bindIP, Port: rpcPort},
		VerboseMode: viper.GetBool(flagVerbose),
	}

	log.Debug().Msgf("advertise IP: %v", advertiseIP.String())
	log.Debug().Msgf("cluster addr: %v", config.ClusterAddr.String())
	log.Debug().Msgf("data dir: %v", dataDir)
	log.Debug().Msgf("http addr: %v", config.HttpAddr.String())
	log.Debug().Msgf("rpc addr: %v", config.RpcAddr.String())

	return controller.StartController(config)
}
