package configs

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type ApiKeyConfig struct {
	Key       string `json:"key"`
	RateLimit int    `json:"rate_limit"`
}

type Conf struct {
	WebServerPort    string `mapstructure:"WEB_SERVER_PORT"`
	IpRateLimit      int    `mapstructure:"IP_RATE_LIMIT"`
	BlockTimeSeconds int    `mapstructure:"BLOCK_TIME_SECONDS"`
	RedisHost        string `mapstructure:"REDIS_HOST"`
	RedisPassword    string `mapstructure:"REDIS_PASSWORD"`
	RedisDB          int    `mapstructure:"REDIS_DB"`
	StorageType      string `mapstructure:"STORAGE_TYPE"`
	ApiKeysRateLimit []ApiKeyConfig
}

func LoadConfig(path string) (*Conf, error) {
	cfg := &Conf{}
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.BindEnv("WEB_SERVER_PORT")
	viper.BindEnv("IP_RATE_LIMIT")
	viper.BindEnv("BLOCK_TIME_SECONDS")
	viper.BindEnv("REDIS_HOST")
	viper.BindEnv("REDIS_PASSWORD")
	viper.BindEnv("REDIS_DB")
	viper.BindEnv("STORAGE_TYPE")
	viper.BindEnv("API_KEYS_RATE_LIMIT")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println(".env file not found, using environment variables")
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("error unmarshalling config: %v", err)
	}

	cfg.ApiKeysRateLimit = []ApiKeyConfig{}

	apiKeysJson := viper.GetString("API_KEYS_RATE_LIMIT")
	if apiKeysJson != "" {
		var apiKeys []ApiKeyConfig
		err = json.Unmarshal([]byte(apiKeysJson), &apiKeys)
		if err != nil {
			log.Fatalf("error parsing API keys JSON: %v", err)
		}
		cfg.ApiKeysRateLimit = apiKeys
	}

	return cfg, nil
}
