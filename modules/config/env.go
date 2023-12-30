package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Execute() {

}

func init() {
	keys := []string{}

	err := godotenv.Load(".env")
	if err != nil {
		return
	}

	for _, key := range keys {
		value, exists := os.LookupEnv(key)
		if exists {
			os.Setenv(key, value)
		}
	}
}
