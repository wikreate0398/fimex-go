package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

type Config struct {
	// .env
	Databases Databases
	RabbitMQ  RabbitMQ

	//.yaml
	Server `mapstructure:"server"`
}

type Databases struct {
	MySql MySql
}

type MySql struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Database string `mapstructure:"DB_DATABASE"`
}

type RabbitMQ struct {
	Host     string
	Port     int
	User     string
	Password string
}

type Server struct {
	Port int `mapstructure:"port"`
}

func Init(env string) (*Config, error) {
	cfg := &Config{}

	if err := cfg.parseYaml(env); err != nil {
		return nil, err
	}

	if err := cfg.parseEnv(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *Config) parseYaml(env string) error {
	envViperInstance := viper.New()

	envViperInstance.AddConfigPath("configs")
	envViperInstance.SetConfigName("main")

	if err := envViperInstance.ReadInConfig(); err != nil {
		return err
	}

	envViperInstance.SetConfigName(env)

	if err := envViperInstance.MergeInConfig(); err != nil {
		return err
	}

	if err := envViperInstance.UnmarshalKey("server", &cfg.Server); err != nil {
		return err
	}

	return nil
}

func (cfg *Config) parseEnv() error {
	if err := gotenv.Load(".env"); err != nil {
		return err
	}

	if err := loadEnv("db", &cfg.Databases.MySql); err != nil {
		return err
	}

	if err := loadEnv("rmq", &cfg.RabbitMQ); err != nil {
		return err
	}

	return nil
}

func loadEnv(prefix string, spec interface{}) error {
	if err := envconfig.Process(prefix, spec); err != nil {
		return err
	}

	return nil
}
