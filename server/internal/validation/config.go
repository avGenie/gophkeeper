package validation

import (
	"fmt"
	"os"
)

// ValidateConfigPath Validates that the path provided is a file, that can be read
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("couldn't get fileinfo: %w", err)
	}

	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}

	return nil
}
