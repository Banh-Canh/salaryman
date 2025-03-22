package fs

import (
	"log/slog"
	"os"

	"github.com/Banh-Canh/salaryman/internal/utils/logger"
)

// EnsureDir ensures directory exits and creates it
func EnsureDir(directory string) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.Mkdir(directory, os.ModePerm)
		if err != nil {
			logger.Logger.Error("EnsureDir - failed to create directory", slog.Any("error", err))
		}
	}
}
