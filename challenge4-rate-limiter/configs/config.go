package configs

import "github.com/spf13/viper"

type Conf struct {
	WebServerPort    string `mapstructure:"WEB_SERVER_PORT"`
	IpRateLimit      int    `mapstructure:"IP_RATE_LIMIT"`
	ApiKeyRateLimit  int    `mapstructure:"API_KEY_RATE_LIMIT"`
	BlockTimeSeconds int    `mapstructure:"BLOCK_TIME_SECONDS"`
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

	return cfg, err
}
