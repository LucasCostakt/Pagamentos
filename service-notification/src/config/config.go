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
	C.AddConfigPath(".")                     // optionally look for config in the working directory
	C.AddConfigPath("/config/")              // optionally look for config in the working directory
	C.SetConfigName("config." + environment) // name of config file (without extension)
	C.AutomaticEnv()                         // read in environment variables that match

	if err := C.ReadInConfig(); err != nil {
		log.Fatal(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
