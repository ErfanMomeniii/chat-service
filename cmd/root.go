package cmd

import (
	"github.com/ErfanMomeniii/chat-service/internal/config"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:                "chat",
		Short:              "chat service",
		Long:               `chat service`,
		PersistentPreRun:   preRun,
		PersistentPostRunE: postRun,
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
}

func preRun(_ *cobra.Command, _ []string) {
	config.Init()
}

func postRun(_ *cobra.Command, _ []string) error {
	return nil
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
