package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// LoadEnv loads the given env file (e.g., ".env", ".env.test") from project root
func LoadEnv(filename string) {
	rootPath, err := getProjectRoot()
	if err != nil {
		log.Printf("Unable to determine project root: %v", err)
		log.Println("Falling back to system environment variables only")
		return
	}

	envPath := filepath.Join(rootPath, filename)
	if err := godotenv.Load(envPath); err != nil {
		log.Printf("%s file not found at: %s, relying on system environment variables", filename, envPath)
	} else {
		log.Printf("Loaded %s from: %s", filename, envPath)
	}
}

func getProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "", os.ErrNotExist
}
