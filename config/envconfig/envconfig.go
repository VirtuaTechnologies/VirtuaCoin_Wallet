package envconfig

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	APP_PORT                       int      `env:"APP_PORT,notEmpty"`
	GIN_MODE                       string   `env:"GIN_MODE,notEmpty"`
	REDIS_HOST                     string   `env:"REDIS_HOST,notEmpty"`
	REDIS_PORT                     int      `env:"REDIS_PORT,notEmpty"`
	REDIS_PASSWORD                 string   `env:"REDIS_PASSWORD,notEmpty"`
	REDIS_DB                       string   `env:"REDIS_DB,notEmpty"`
	ALLOWED_ORIGIN                 []string `env:"ALLOWED_ORIGIN,notEmpty" envSeparator:","`
	NETWORK_RPC_URL_ETHEREUM       string   `env:"NETWORK_RPC_URL_ETHEREUM,notEmpty"`
	NETWORK_RPC_URL_POLYGON        string   `env:"NETWORK_RPC_URL_POLYGON,notEmpty"`
	NETWORK_RPC_URL_BSC            string   `env:"NETWORK_RPC_URL_BSC,notEmpty"`
	TOKEN                          string   `env:"TOKEN,notEmpty"`
	APP_ENVIRONMENT                string   `env:"APP_ENVIRONMENT,notEmpty"`
	OPERATOR_MNEMONIC              string   `env:"OPERATOR_MNEMONIC,notEmpty"`
	VIRTUACOINNFT_CONTRACT_ADDRESS string   `env:"VIRTUACOINNFT_CONTRACT_ADDRESS,notEmpty"`
	CHAIN_ID_POLYGON               int      `env:"CHAIN_ID_POLYGON,notEmpty"`
	CHAIN_ID_BSC                   int      `env:"CHAIN_ID_BSC,notEmpty"`
	CHAIN_ID_ETHEREUM              int      `env:"CHAIN_ID_ETHEREUM,notEmpty"`
}

var EnvVars config = config{}

func InitEnvVars() {
	if err := env.Parse(&EnvVars); err != nil {
		log.Fatalf("failed to parse EnvVars: %s", err)
	}
	gin.SetMode(EnvVars.GIN_MODE)
}
