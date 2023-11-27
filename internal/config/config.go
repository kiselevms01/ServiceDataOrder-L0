package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/kiselevms01/wbProject_L0/internal/http"
	"github.com/nats-io/stan.go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Database      DatabaseConfig      `yaml:"database"`
	Http          http.Config         `yaml:"http"`
	NatsStreaming NatsStreamingConfig `yaml:"nats_streaming"`
}

type NatsStreamingConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	ClusterID string `yaml:"cluster_id"`
	ClientID  string `yaml:"client_id"`
	Channel   string `yaml:"channel"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}

func MustLoad() *Config {
	os.Setenv("CONFIG_PATH", "config.local.yaml")
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// check if file exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}

func DbConnect(cfg DatabaseConfig) (*gorm.DB, error) {
	dbData := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dbData), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, err
}

func NatsStreamingConnect(cfg NatsStreamingConfig) (stan.Conn, error) {
	host := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	sc, err := stan.Connect(cfg.ClusterID+"", cfg.ClientID, stan.NatsURL(host))
	if err != nil {
		return nil, err
	}

	return sc, err
}
