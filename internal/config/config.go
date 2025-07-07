package config

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	HTTPServer `mapstructure:"http_server"`
	DBConfig   `mapstructure:"db"`
}

type HTTPServer struct {
	Address      string        `mapstructure:"address"`
	Port         int           `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"read_t"`
	WriteTimeout time.Duration `mapstructure:"write_t"`
}

type DBConfig struct {
	Address  string `mapstructure:"address"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

func SetupConfig() *Config {
	path := os.Getenv("CONFIG_PATH")
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic("failed read config file")
	}
	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic("failed read config file")
	}
	return &cfg
}
