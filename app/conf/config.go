package conf

import (
	"fmt"

	"github.com/jinzhu/configor"
)

// Config is
var Config = struct {
	APP struct {
		Name string `default:"app name"`
		Port string `default:"18000"`
	}

	DB struct {
		Name     string
		Host     string `default:"localhost"`
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     string `default:"3306"`
		Driver   string ``
	}

	Contacts []struct {
		Name  string
		Email string `required:"true"`
	}
}{}

func init() {
	configor.Load(&Config, "config.yml")
	fmt.Printf("config: %#v", Config)
}
