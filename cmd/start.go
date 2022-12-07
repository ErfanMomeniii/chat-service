package cmd

import (
	"github.com/ErfanMomeniii/chat-service/internal/app"
	internalhttp "github.com/ErfanMomeniii/chat-service/internal/http"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Translator service.",
	Long:  `Start the Translator service.`,
	Run:   startFunc,
}

func startFunc(cmd *cobra.Command, args []string) {
	app.WithGracefulShutdown()

	internalhttp.
		NewServer().
		Serve()

	app.Wait()
}
