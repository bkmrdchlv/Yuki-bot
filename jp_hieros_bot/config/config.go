package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment string
	BotToken    string
	LogLevel    string
}

func Load() Config {
	// godotenv.Load(".env")
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "development"))
	c.BotToken = cast.ToString(getOrReturnDefault("TELEGRAMM_BOT_TOKEN", "7562383282:AAG2Fp7oijMNDvBYjL2Z9_hZLXke1LNzats"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	return c
}
func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
