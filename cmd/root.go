package cmd

import (
	"github.com/ErfanMomeniii/chat-service/internal/config"
	"github.com/ErfanMomeniii/chat-service/internal/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
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
	err := log.Level.UnmarshalText([]byte(config.C.Logger.Level))
	if err != nil {
		log.Logger.With(zap.Error(err)).Fatal("error in setting log level from config")
	}
}

func postRun(_ *cobra.Command, _ []string) error {
	return log.CloseLogger()
}

func Execute() error {
	return rootCmd.Execute()
}
