package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Build   string
	Version string
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Joat - v%s (build: %s)\n", Version, Build)
	},
}
