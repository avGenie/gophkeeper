package logger

import (
	"go.uber.org/zap"
)

// Initialize Initializes zap logger
func Initialize(level, outputPath string) error {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return err
	}

	cfg := zap.NewProductionConfig()
	cfg.Level = lvl

	if outputPath != "" {
		cfg.ErrorOutputPaths = []string{outputPath}
		cfg.OutputPaths = []string{outputPath}
	}

	zl, err := cfg.Build()
	if err != nil {
		return err
	}

	zap.ReplaceGlobals(zl)

	return nil
}
