package logger

import (
	"github.com/avGenie/gophkeeper/server/internal/app/config"
	"go.uber.org/zap"
)

// Initialize Initializes zap logger
func Initialize(config config.Config) error {
	lvl, err := zap.ParseAtomicLevel(config.LogLevel)
	if err != nil {
		return err
	}

	cfg := zap.NewProductionConfig()
	cfg.Level = lvl

	zl, err := cfg.Build()
	if err != nil {
		return err
	}

	zap.ReplaceGlobals(zl)

	return nil
}
