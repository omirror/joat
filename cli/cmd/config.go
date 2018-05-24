package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	//controllerCmd.Flags().String(flagAdvertiseAddr, "", flagAdvertiseAddrDesc)
	//controllerCmd.Flags().String(flagBindAddr, defaultBindAddr, flagBindAddrDesc)
	//controllerCmd.Flags().Int(flagClusterPort, defaultClusterPort, flagClusterPortDesc)
	//controllerCmd.Flags().String(flagDataDir, defaultDataDir, flagDataDirDesc)
	//controllerCmd.Flags().String(flagHttpBindAddr, defaultBindAddr, flagHttpBindAddrDesc)
	//controllerCmd.Flags().Int(flagHttpPort, defaultHttpPort, flagHttpPortDesc)
	//controllerCmd.Flags().Int(flagRpcPort, defaultRpcPort, flagRpcPortDesc)
	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use: "config",
}
