package config

import "github.com/spf13/viper"

var Cfg *viper.Viper

type Config struct {
	Application *System `yaml:"system"`
	Logger      *Logger `yaml:"logger"`
}

func SetUpConfig(path string) {
	Cfg := viper.New()
	Cfg.SetConfigFile(path)
	Cfg.SetConfigType("yaml")
	if err := Cfg.ReadInConfig(); err != nil {
		panic(err)
	}
}
