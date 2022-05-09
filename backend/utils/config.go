package utils

import "github.com/spf13/viper"

type Config struct {
	DISCORD_TOKEN string `mapstructure:"DISCORD_TOKEN"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}