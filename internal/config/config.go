package config

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
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
	viper.SetDefault("http_server.address", "localhost")
	viper.SetDefault("http_server.port", 8081)
	viper.SetDefault("http_server.read_t", "5s")
	viper.SetDefault("http_server.write_t", "5s")

	viper.SetDefault("db.address", "localhost")
	viper.SetDefault("db.port", 5433)
	viper.SetDefault("db.user", "postgres")
	viper.SetDefault("db.password", "12332187")
	viper.SetDefault("db.name", "postgres")

	path := os.Getenv("CONFIG_PATH")
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Error("error reading configuration, default values will be used")
		return &Config{
			HTTPServer: HTTPServer{},
			DBConfig:   DBConfig{},
		}
	}
	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		logrus.Error("error unmarshal configuration, default values will be used")
		return &Config{
			HTTPServer: HTTPServer{},
			DBConfig:   DBConfig{},
		}
	}
	return &cfg
}
