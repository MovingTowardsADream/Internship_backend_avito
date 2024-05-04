package configs

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	App      AppConfig      `yaml:"app"`
	Database DatabaseConfig `yaml:"db"`
}

type AppConfig struct {
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	Username string `yaml:"username" env-required:"true"`
	Host     string `yaml:"host" env-default:"local"`
	Port     string `yaml:"port" env-required:"true"`
	DBName   string `yaml:"dbname" env-required:"true"`
	SSLMode  string `yaml:"sslmode" env-required:"true"`
}

func MustLoadByPath(configPath string) *Config {
	// Check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("Config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("Cannot read config: " + err.Error())
	}

	return &cfg
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("Config path is empty")
	}

	return MustLoadByPath(path)
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "./configs/config.yaml", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
