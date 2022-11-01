package config

import "github.com/spf13/viper"

var Cfg *viper.Viper

type Config struct {
	Application *System `yaml:"system"`
	Logger      *Logger `yaml:"logger"`
}

func SetUpConfig(path string) {
	cfg := viper.New()
	cfg.SetConfigFile(path)
	cfg.SetConfigType("yaml")
	if err := cfg.ReadInConfig(); err != nil {
		panic(err)
	}

}
