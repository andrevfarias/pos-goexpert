package configs

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

type ApiKeyConfig struct {
	Key       string `json:"key"`
	RateLimit int    `json:"rate_limit"`
}

type Conf struct {
	WebServerPort    string `mapstructure:"WEB_SERVER_PORT"`
	IpRateLimit      int    `mapstructure:"IP_RATE_LIMIT"`
	ApiKeyRateLimit  int    `mapstructure:"API_KEY_RATE_LIMIT"`
	BlockTimeSeconds int    `mapstructure:"BLOCK_TIME_SECONDS"`
	RedisHost        string `mapstructure:"REDIS_HOST"`
	RedisPassword    string `mapstructure:"REDIS_PASSWORD"`
	RedisDB          int    `mapstructure:"REDIS_DB"`
	ApiKeysRateLimit []ApiKeyConfig
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	cfg.ApiKeysRateLimit = []ApiKeyConfig{}

	apiKeysJson := viper.GetString("API_KEYS_RATE_LIMIT")
	if apiKeysJson != "" {
		var apiKeys []ApiKeyConfig
		err = json.Unmarshal([]byte(apiKeysJson), &apiKeys)
		if err != nil {
			panic(fmt.Errorf("error parsing API keys JSON: %w", err))
		}
		cfg.ApiKeysRateLimit = apiKeys
	}

	return cfg, nil
}
