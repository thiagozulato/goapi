package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	goEnv = os.Getenv("GO_ENV")
)

// Get recupera configuraçoes a partir de variaveis de ambiente
func Get(key string) string {
	if goEnv == "" {
		goEnv = "development"
	}

	if err := godotenv.Load(fmt.Sprintf(".env.%s", goEnv)); err != nil {
		log.Panic("Variáveis de ambiente não foram carregadas")
	}

	return os.Getenv(key)
}
