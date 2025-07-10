package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// type HTTPServer struct {
// 	Adds string `yaml:"address" env-required:"true"`
// }

// env - default:""production
//
//	type Config struct {
//		Env         string `yaml:"env" env:"ENV" env-rrequired:"true"`
//		StoragePath string `yaml:"storage_path" env-required:"true"`
//		HTTPServer  `yaml:"http_server"`
//	}
type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type HTTPServer struct {
	Address string `yaml:"address"`
}

type ConfigPSQL struct {
	Env        string         `yaml:"env"`
	Postgres   PostgresConfig `yaml:"postgres"`
	HTTPServer HTTPServer     `yaml:"http_server"`
}

// changes Config to  ConfigPSQL
func MustLoad() *ConfigPSQL {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {

		flags := flag.String("config", "", "path to the configuration file ")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set")
		}

	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exits: %s ", configPath)
	}
	// Change it
	var cfg ConfigPSQL

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("can not read config file %s ", err.Error())

	}

	return &cfg

}
