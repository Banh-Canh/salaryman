package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"

	"github.com/Banh-Canh/salaryman/internal/utils/logger"
)

var RootCmd = &cobra.Command{
	Use:   "salaryman",
	Short: "salaryman",
	Long: `
Build your resume with HTML/CSS and JSON Data
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Initialize configuration here
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help() //nolint:all
	},
}

func Execute() {
	logger.InitializeLogger(slog.LevelInfo)
	if err := RootCmd.Execute(); err != nil {
		logger.Logger.Error("error", slog.Any("error", err))
	}
}
