package server

import (
	"log/slog"

	"github.com/spf13/cobra"

	"github.com/Banh-Canh/salaryman/internal/api"
	"github.com/Banh-Canh/salaryman/internal/utils/logger"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts a server to use app as an API",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.InitializeLogger(slog.LevelInfo)
		if err := api.New().Run(); err != nil {
			return err
		}
		return nil
	},
}

func Cmd() *cobra.Command {
	return serverCmd
}
