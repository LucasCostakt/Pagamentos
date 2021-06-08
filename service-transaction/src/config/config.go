package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var C *viper.Viper

func init() {
	C = viper.New()
}

var Env string

func LoadConfig(environment string) {
	if environment == "" {
		environment = "development"
		Env = environment
	}

	C.SetConfigType("toml")
	C.AddConfigPath(".")
	C.AddConfigPath("/config/")
	C.SetConfigName("config." + environment)
	C.AutomaticEnv()

	if err := C.ReadInConfig(); err != nil {
		log.Fatal(fmt.Errorf("fatal error config file: %w", err))
	}
}
