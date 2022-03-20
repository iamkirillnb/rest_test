package internal

import (
	"github.com/iamkirillnb/rus_law/internal/entities"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type ServerConfig struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Config struct {
	DbConfig entities.DbRepo `yaml:"postgres_local"`
	Server   ServerConfig    `yaml:"server_http"`
}

var instance *Config
var once sync.Once

func GetConfig(path string) *Config {
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadConfig(path, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			log.Fatal(err)
		}
	})
	return instance
}
