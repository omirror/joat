package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	configCmd.AddCommand(configSetCmd)
}

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set configuration property",
	RunE:  configSetRun,
}

func configSetRun(_ *cobra.Command, _ []string) error {
	//log.Info().Msg("Starting controller instance...")
	//
	//advertiseAddr := viper.GetString(flagAdvertiseAddr)
	//if advertiseAddr == "" {
	//	resolvedAdvertiseAddr, err := template.Parse("{{ GetPrivateIP }}")
	//	if err != nil {
	//		return fmt.Errorf("failed to get advertise IP address: %v", err)
	//	}
	//	advertiseAddr = resolvedAdvertiseAddr
	//}
	//bindAddr := viper.GetString(flagBindAddr)
	//httpBindAddr := viper.GetString(flagHttpBindAddr)
	//
	//advertiseIP := net.ParseIP(advertiseAddr)
	//bindIP := net.ParseIP(bindAddr)
	//httpServerIP := net.ParseIP(httpBindAddr)
	//
	//clusterPort := viper.GetInt(flagClusterPort)
	//if !util.IsValidPort(clusterPort) {
	//	return fmt.Errorf("invalid cluster port: %d", clusterPort)
	//}
	//
	//httpPort := viper.GetInt(flagHttpPort)
	//if !util.IsValidPort(httpPort) {
	//	return fmt.Errorf("invalid http port: %d", httpPort)
	//}
	//
	//rpcPort := viper.GetInt(flagRpcPort)
	//if !util.IsValidPort(rpcPort) {
	//	return fmt.Errorf("invalid rpc port: %d", rpcPort)
	//}
	//
	//dataDir, _ := homedir.Expand(viper.GetString(flagDataDir))
	//
	//config := &controller.Config{
	//	AdvertiseIP: &advertiseIP,
	//	ClusterAddr: &net.TCPAddr{IP: bindIP, Port: clusterPort},
	//	DataDir:     dataDir,
	//	HttpAddr:    &net.TCPAddr{IP: httpServerIP, Port: httpPort},
	//	RpcAddr:     &net.TCPAddr{IP: bindIP, Port: rpcPort},
	//	VerboseMode: viper.GetBool(flagVerbose),
	//}
	//
	//log.Debug().Msgf("advertise IP: %v", advertiseIP.String())
	//log.Debug().Msgf("cluster addr: %v", config.ClusterAddr.String())
	//log.Debug().Msgf("data dir: %v", dataDir)
	//log.Debug().Msgf("http addr: %v", config.HttpAddr.String())
	//log.Debug().Msgf("rpc addr: %v", config.RpcAddr.String())

	return nil
}
