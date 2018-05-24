package cmd

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:               "joat",
	PersistentPreRunE: rootPreRun,
}

func Execute() error {
	rootCmd.PersistentFlags().BoolP(flagDebug, flagDebugShort, false, flagDebugDesc)
	rootCmd.PersistentFlags().BoolP(flagVerbose, flagVerboseShort, false, flagVerboseDesc)
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func rootPreRun(cmd *cobra.Command, _ []string) error {
	// Configure viper
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	viper.SetEnvPrefix(configEnvPrefix)
	viper.SetEnvKeyReplacer(envKeyReplacer)
	viper.AutomaticEnv()

	// Enable debug log if verbose flag is set
	if viper.GetBool(flagVerbose) {
		viper.Set(flagDebug, true)
	}

	// Configure logger
	if viper.GetBool(flagDebug) {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	return nil
}
