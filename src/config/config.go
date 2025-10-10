package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Password PasswordConfig
	Logger   LoggingConfig
}

type ServerConfig struct {
	Port    string
	RunMode string
}

type LoggingConfig struct {
	Filepath string
	Encoding string
	Level    string
	Logger   string
}

type PostgresConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DbName          string
	SSLMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type RedisConfig struct {
	Host              string
	Port              string
	Password          string
	Db                string
	DialTimeout       time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	PoolSize          int
	PoolTimeout       time.Duration
	IdlCheckFrequency time.Duration
}

type PasswordConfig struct {
	includeChars     bool
	includeDigits    bool
	minLength        int
	maxLength        int
	includeUppercase bool
	includeLowercase bool
}

func GetConfig() *Config {
	configPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(configPath, "yml")
	if err != nil {
		log.Fatalf("error in loading config %v \n", err)
	}

	cfg, err := ParseConfig(v)
	if err != nil {
		log.Fatalf("error in parsing config %v \n", err)
	}

	return cfg
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var cnfg Config
	err := v.Unmarshal(&cnfg)
	if err != nil {
		return nil, err
	}

	return &cnfg, nil
}

func LoadConfig(fileName string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigFile(fileName)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("file not found")
		}
		return nil, err
	}
	return v, nil
}

func getConfigPath(env string) string {
	switch env {
	case "docker":
		return "../config/config-docker.yml"
	case "production":
		return "config/config-production.yml"
	default:
		return "../config/config-dev.yml"
	}

}
