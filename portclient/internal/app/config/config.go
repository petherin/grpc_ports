package config

import (
	"os"
)

const (
	FILE_PATH       = "FILE_PATH"
	SVC_URL         = "SVC_URL"
	DEFAULT_PATH    = "files/ports.json" // Default file path that works locally unless given alternative.
	DEFAULT_SVC_URL = "0.0.0.0:50051"    // Default service address that works locally unless given alternative.
)

type Config struct {
	FilePath string
	SvcURL   string
	Port     string
}

// Get retrieves config held in environment variables
func Get() (Config, error) {
	filePath := os.Getenv(FILE_PATH)
	if len(filePath) == 0 {
		filePath = DEFAULT_PATH
	}

	svcURL := os.Getenv(SVC_URL)
	if len(svcURL) == 0 {
		svcURL = DEFAULT_SVC_URL
	}

	return Config{
		FilePath: filePath,
		SvcURL:   svcURL,
	}, nil
}
