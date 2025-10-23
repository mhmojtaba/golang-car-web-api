package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Password PasswordConfig
	Logger   LoggingConfig
	Otp      OtpConfig
	Jwt      JwtConfig
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
	IncludeChars     bool
	IncludeDigits    bool
	MinLength        int
	MaxLength        int
	IncludeUppercase bool
	IncludeLowercase bool
}

type OtpConfig struct {
	ExpireTime time.Duration
	Digits     int
	Limiter    time.Duration
}

type JwtConfig struct {
	SecretKey                  string
	RefreshSecretKey           string
	AccessTokenExpireDuration  time.Duration
	RefreshTokenExpireDuration time.Duration
}

func GetConfig() *Config {
	env := os.Getenv("APP_ENV")
	configPath := getConfigPath(env)
	log.Printf("Loading config for env: %s from path: %s\n", env, configPath)

	v, err := LoadConfig(configPath, "yml")
	if err != nil {
		log.Fatalf("error in loading config from path %s: %v\n", configPath, err)
	}

	cfg, err := ParseConfig(v)
	if err != nil {
		log.Fatalf("error in parsing config: %v\n", err)
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
	fullPath := fileName + "." + fileType
	log.Printf("Attempting to load config from: %s", fullPath)

	absPath, err := filepath.Abs(fullPath)
	if err != nil {
		return nil, fmt.Errorf("error getting absolute path for %s: %v", fullPath, err)
	}
	log.Printf("Absolute config path: %s", absPath)

	v.SetConfigFile(fullPath)

	// Check if file exists before trying to read it
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found at %s", fullPath)
	}

	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	log.Printf("Successfully loaded config from: %s", fullPath)
	return v, nil
}

func getConfigPath(env string) string {
	if env == "docker" {
		return "/app/config/config-docker"
	} else if env == "production" {
		return "/config/config-production"
	} else {
		return "config/config-development"
	}
}
