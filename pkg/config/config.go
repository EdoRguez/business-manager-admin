package config

import "github.com/spf13/viper"

type Config struct {
	Db_Company string `mapstructure:"DB_COMPANY"`
	Db_Auth    string `mapstructure:"DB_AUTH"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
