package main

import (
	"log/slog"

	"github.com/spf13/cobra"

	"github.com/Banh-Canh/salaryman/cmd/local"
	"github.com/Banh-Canh/salaryman/cmd/server"
	"github.com/Banh-Canh/salaryman/cmd/version"
	"github.com/Banh-Canh/salaryman/internal/utils/logger"
)

var rootCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(local.Cmd())
	rootCmd.AddCommand(server.Cmd())
	rootCmd.AddCommand(version.Cmd())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logger.Logger.Error("error", slog.Any("error", err))
	}
}
