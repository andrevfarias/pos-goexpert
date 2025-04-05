package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Config representa as configurações da aplicação
type Config struct {
	Port              string `mapstructure:"PORT"`
	ViacepAPIBaseURL  string `mapstructure:"VIACEP_API_BASE_URL"`
	WeatherAPIBaseURL string `mapstructure:"WEATHER_API_BASE_URL"`
	WeatherAPIKey     string `mapstructure:"WEATHER_API_KEY"`
	APITimeoutSeconds int    `mapstructure:"API_TIMEOUT_SECONDS"`
}

// LoadConfig carrega as configurações do arquivo .env e variáveis de ambiente
func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.BindEnv("PORT")
	viper.BindEnv("VIACEP_API_BASE_URL")
	viper.BindEnv("WEATHER_API_BASE_URL")
	viper.BindEnv("WEATHER_API_KEY")
	viper.BindEnv("API_TIMEOUT_SECONDS")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("erro ao ler arquivo de configuração: %w", err)
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("erro ao decodificar configurações: %w", err)
	}

	// Validações
	if config.WeatherAPIKey == "" {
		return nil, fmt.Errorf("WEATHER_API_KEY é obrigatória")
	}

	return &config, nil
}
