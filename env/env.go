package env

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	cenv "github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	Port         string = "3030"
	IsProduction bool   = false
	TemplateDir         = "templates"
)

type EnvConfig struct {
	IsProduction bool   `env:"PRODUCTION" envDefault:"false"`
	Port         string `env:"PORT" envDefault:"3030"`
}

var EnvVars = &EnvConfig{}

func LoadEnv() error {
	var err error
	err = cenv.Parse(EnvVars)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", EnvVars)

	Port = EnvVars.Port
	IsProduction = EnvVars.IsProduction

	if EnvVars.IsProduction {

		log.Println("[production]")

		gin.SetMode(gin.ReleaseMode)
		os.Setenv("GIN_MODE", "release")

	} else {
		filename := ".env.development"
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			filename = ".env"
		}
		err = godotenv.Load(filename)
		if err != nil {
			log.Warn(".env file not loaded")
		}
	}
	return err
}
