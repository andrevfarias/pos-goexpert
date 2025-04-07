package config

import (
	"fmt"
	"path/filepath"
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
	envPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter caminho absoluto: %w", err)
	}

	viper.AddConfigPath(envPath)
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.SetConfigFile(filepath.Join(envPath, ".env"))

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

	// Definindo valores padrão
	if config.Port == "" {
		config.Port = "8080" // Valor padrão para a porta
	}
	if config.ViacepAPIBaseURL == "" {
		config.ViacepAPIBaseURL = "https://viacep.com.br/ws" // Valor padrão para a API do ViaCEP
	}
	if config.WeatherAPIBaseURL == "" {
		config.WeatherAPIBaseURL = "https://api.weatherapi.com/v1" // Valor padrão para a API do clima
	}

	// Validações
	if config.WeatherAPIKey == "" {
		return nil, fmt.Errorf("WEATHER_API_KEY é obrigatória")
	}

	return &config, nil
}
