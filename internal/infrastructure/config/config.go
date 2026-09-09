package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	App      AppConfig `toml:"app"`
	Database DBConfig  `toml:"database"`
	Log      LogConfig `toml:"log"`
}

type AppConfig struct {
	Name            string   `toml:"name"`
	Port            int64    `toml:"port"`
	ShutdownTimeout int      `toml:"shutdown_timeout"` // in seconds
	AllowedOrigins  []string `toml:"allowed_origins"`
}

type DBConfig struct {
	URL             string `toml:"url"`
	MaxConns        int    `toml:"max_conns"`
	MinConns        int    `toml:"min_conns"`
	MaxConnLifetime int    `toml:"max_conn_lifetime"`  // in minutes
	MaxConnIdleTime int    `toml:"max_conn_idle_time"` // in minutes
}

type LogConfig struct {
	Level      string `toml:"level"`
	FilePath   string `toml:"file_path"`
	MaxSize    int    `toml:"max_size"`    // in megabytes
	MaxBackups int    `toml:"max_backups"` // number of backups
	MaxAge     int    `toml:"max_age"`     // in days
}

// Load loads configuration from a TOML file
func Load(path string) (*Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found: %s", path)
	}

	var cfg Config
	// Set default values
	cfg.Database.MaxConns = 10
	cfg.Database.MinConns = 2
	cfg.Database.MaxConnLifetime = 30
	cfg.Database.MaxConnIdleTime = 5

	// Set default log values
	cfg.Log.Level = "info"
	cfg.Log.FilePath = "logs/app.log"
	cfg.Log.MaxSize = 100
	cfg.Log.MaxBackups = 3
	cfg.Log.MaxAge = 28

	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
