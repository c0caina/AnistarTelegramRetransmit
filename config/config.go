package config

import "github.com/spf13/viper"

type Config struct{ viper *viper.Viper }

func NewConfig() (*Config, error) {
	runtimeViper := viper.New()
	runtimeViper.AddConfigPath("config/.")
	runtimeViper.SetConfigName("cfg")
	runtimeViper.SetConfigType("toml")

	err := runtimeViper.ReadInConfig()

	return &Config{viper: runtimeViper}, err
}

func (c Config) CheckCommand(cmd string) string {
	return c.viper.GetString(cmd)
}
