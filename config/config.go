package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	AppName              string
	AppEnv               string
	RedisAddr            string
	DBUserName           string
	DBPassword           string
	DBHostWriter         string
	DBHostReader         string
	DBPort               string
	DBName               string
	DBMaxOpenConnections int
	DBMaxIdleConnections int
	ServerPort           string
	JWT_SECRET           string
}

var config Config

// Should run at the very beginning, before any other package
// or code.
func init() {
	appEnv := os.Getenv("APP_ENV")
	if len(appEnv) == 0 {
		appEnv = "dev"
	}

	configFilePath := "./config/.env"

	switch appEnv {
	case "production":
		configFilePath = "./config/.env.prod"
	case "stage":
		configFilePath = "./config/.env.stage"
	}
	zap.L().Info("Loading env file", zap.String("path", configFilePath))

	e := godotenv.Load(configFilePath)
	if e != nil {
		zap.L().Error("Error loading .env file", zap.Error(e))
		return
	}
	config.AppName = os.Getenv("SERVICE_NAME")
	config.AppEnv = appEnv
	config.RedisAddr = os.Getenv("REDIS_ADDR")
	config.DBUserName = os.Getenv("DB_USERNAME")
	config.DBHostReader = os.Getenv("DB_HOST_READER")
	config.DBHostWriter = os.Getenv("DB_HOST_WRITER")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBName = os.Getenv("DB_NAME")
	config.DBMaxIdleConnections, _ = strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONENCTION"))
	config.DBMaxOpenConnections, _ = strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
	config.ServerPort = os.Getenv("SERVER_PORT")
	config.JWT_SECRET = os.Getenv("JWT_SECRET")
}

func Get() Config {
	return config
}

func IsProduction() bool {
	return config.AppEnv == "production"
}
