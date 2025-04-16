package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port int `mapstructure:"Port"`
	DB   struct {
		DNS          string `mapstructure:"Dns"`
		MaxOpenConns int    `mapstructure:"MaxOpenConns"`
		MaxIdleConns int    `mapstructure:"MaxIdleConns"`
		MaxIdleTime  string `mapstructure:"MaxIdleTime"`
	} `mapstructure:"Database"`
	Redis struct {
		Port     int    `mapstructure:"Port"`
		Host     string `mapstructure:"Host"`
		DB       int    `mapstructure:"DB"`
		Password string `mapstructure:"Password"`
	} `mapstructure:"Redis"`
	Kafka struct {
		Brokers    []string `mapstructure:"Brokers"`
		Topic      string   `mapstructure:"Topic"`
		TopicGroup string   `mapstructure:"TopicGroup"`
	} `mapstructure:"Kafka"`
}

func LoadConfig() (*Config, error) {
	var cfg *Config
	v := viper.New()

	v.AddConfigPath(".")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	err := v.Unmarshal(&cfg)

	return cfg, err
}
