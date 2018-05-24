package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initializeCmd)
}

var initializeCmd = &cobra.Command{
	Use:   "initialize <controller_address>",
	Short: "Initialize a new controller instance",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("required argument: controller_address")
		}
		return nil
	},
	RunE: initializeRun,
}

func initializeRun(_ *cobra.Command, _ []string) error {
	return nil
}
