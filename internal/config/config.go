package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv      string
	DBURL       string
	RedisAddr   string
	KafkaBroker string
	KafkaTopic  string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		AppEnv:      os.Getenv("APP_ENV"),
		DBURL:       os.Getenv("DB_URL"),
		RedisAddr:   os.Getenv("REDIS_ADDR"),
		KafkaBroker: os.Getenv("KAFKA_BROKER"),
		KafkaTopic:  os.Getenv("KAFKA_TOPIC"),
	}

	if cfg.DBURL == "" {
		log.Fatal("DB_URL не задан")
	}

	return cfg
}
