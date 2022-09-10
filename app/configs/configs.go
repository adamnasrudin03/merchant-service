package configs

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Name      string
	Env       string
	Port      string
	SecretKey string
}

type DbConfig struct {
	Host        string
	Port        string
	Dbname      string
	Username    string
	Password    string
	DbIsMigrate bool
}

type Configs struct {
	Appconfig AppConfig
	Dbconfig  DbConfig
}

var lock = &sync.Mutex{}
var configs *Configs

func GetInstance() *Configs {
	if configs == nil {
		lock.Lock()

		if err := godotenv.Load(); err != nil {
			panic("Failed to load env file")
		}

		configs = &Configs{
			Appconfig: AppConfig{
				Name:      os.Getenv("APP_NAME"),
				Env:       os.Getenv("APP_ENV"),
				Port:      os.Getenv("APP_PORT"),
				SecretKey: os.Getenv("JWT_SECRET"),
			},
			Dbconfig: DbConfig{
				Host:        os.Getenv("DB_HOST"),
				Port:        os.Getenv("DB_PORT"),
				Dbname:      os.Getenv("DB_NAME"),
				Username:    os.Getenv("DB_USER"),
				Password:    os.Getenv("DB_PASS"),
				DbIsMigrate: os.Getenv("DB_ISMIGRATE") == "true",
			},
		}
		lock.Unlock()
	}

	return configs
}
