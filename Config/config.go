package config

import "os"

// Config recupera configuraçoes a partir de variaveis de ambiente
func Config(key string) string {
	port := os.Getenv(key)

	if port == "" {
		port = "5000"
	}

	return port
}
