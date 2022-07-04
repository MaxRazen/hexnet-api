package common

import (
	"github.com/joho/godotenv"
	"os"
)

const EnvServerHost = "SERVER_HOST"
const EnvServerPort = "SERVER_PORT"
const EnvDbConnection = "DB_CONNECTION"
const EnvDbHost = "DB_HOST"
const EnvDbDatabase = "DB_DATABASE"
const EnvDbUsername = "DB_USERNAME"
const EnvDbPassword = "DB_PASSWORD"
const EnvDbPath = "DB_PATH"

var config Config

type DbConfig struct {
	connection string
	host       string
	database   string
	username   string
	password   string
	path       string
}

type AppEnv struct {
	DB         DbConfig
	ServerHost string
	ServerPort string
}

type Config struct {
	Env     AppEnv
	AppName string
}

func LoadConfig(envPath string) Config {
	if envPath == "" {
		envPath = ".env"
	}

	err := godotenv.Load(envPath)

	if err != nil {
		panic("Fatal: Environment configuration file not found in " + envPath)
	}

	config = Config{
		AppName: "hexnet/api",
		Env: AppEnv{
			ServerHost: os.Getenv(EnvServerHost),
			ServerPort: resolveServerPort(),
			DB: DbConfig{
				connection: os.Getenv(EnvDbConnection),
				host:       os.Getenv(EnvDbHost),
				database:   os.Getenv(EnvDbDatabase),
				username:   os.Getenv(EnvDbUsername),
				password:   os.Getenv(EnvDbPassword),
				path:       os.Getenv(EnvDbPath),
			},
		},
	}

	return config
}

func GetConfig() Config {
	return config
}

func resolveServerPort() string {
	port := os.Getenv(EnvServerPort)
	if port == "" {
		return "8080"
	}
	return port
}
