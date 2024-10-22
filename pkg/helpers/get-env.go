package helpers

import (
	"fmt"
	"os"
)

func GetEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		fmt.Printf("No %v environment variable found", name)
		os.Exit(1)
	}

	return value
}
