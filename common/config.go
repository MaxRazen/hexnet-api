package common

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

const (
	EnvServerHost   = "SERVER_HOST"
	EnvServerPort   = "SERVER_PORT"
	EnvDbConnection = "DB_CONNECTION"
	EnvDbHost       = "DB_HOST"
	EnvDbDatabase   = "DB_DATABASE"
	EnvDbUsername   = "DB_USERNAME"
	EnvDbPassword   = "DB_PASSWORD"
	EnvDbPath       = "DB_PATH"
	EnvAuthTimeout  = "AUTH_TIMEOUT"
	EnvAuthSecret   = "AUTH_SECRET"
	EnvAuthRefresh  = "AUTH_REFRESH_TTL"
)

var config Config

type DbConfig struct {
	connection string
	host       string
	database   string
	username   string
	password   string
	path       string
}

type AuthConfig struct {
	Secret  string
	Timeout int
	Refresh int
}

type AppEnv struct {
	DB         DbConfig
	Auth       AuthConfig
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
			ServerPort: resolveStr(os.Getenv(EnvServerPort), "8080"),
			DB: DbConfig{
				connection: os.Getenv(EnvDbConnection),
				host:       os.Getenv(EnvDbHost),
				database:   os.Getenv(EnvDbDatabase),
				username:   os.Getenv(EnvDbUsername),
				password:   os.Getenv(EnvDbPassword),
				path:       os.Getenv(EnvDbPath),
			},
			Auth: AuthConfig{
				Secret:  os.Getenv(EnvAuthSecret),
				Refresh: resolveInt(os.Getenv(EnvAuthRefresh), 60),
				Timeout: resolveInt(os.Getenv(EnvAuthTimeout), 15),
			},
		},
	}

	return config
}

func GetConfig() Config {
	return config
}

func resolveStr(str, def string) string {
	if str == "" {
		return def
	}
	return str
}

func resolveInt(str string, def int) int {
	if str == "" {
		return def
	}
	v, e := strconv.Atoi(str)
	if e != nil {
		return def
	}
	return v
}
