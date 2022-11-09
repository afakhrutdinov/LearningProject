package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	//MaxAttempts int    `json:"maxAttempts"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Printf("read application configuration\n")
		instance = &Config{}
		if err := cleanenv.ReadConfig("pkg/config.yml", instance); err != nil {
			fmt.Println("Err: ", err)
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Fatalf("%v", help)
			//logger.Info(help)
			//logger.Fatal(err)
		}
		fmt.Println("Err ReadConfig 2")
	})
	return instance
}
