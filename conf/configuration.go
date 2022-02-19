package conf

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var Config *Configuration

type Configuration struct {
	Port int `required:"true" envconfig:"PORT"`
	DatabaseURL string `required:"true" envconfig:"DATABASE_URL"`
	FirebaseServiceAccount string `required:"true" split_words:"true"`
}

func loadConfiguration() *Configuration {
	if err := godotenv.Load();err != nil {
		fmt.Print(err)
	}

	conf :=new(Configuration)
	if err := envconfig.Process("JC",conf);err!=nil{
		fmt.Print(err)
	}
	return conf
}

func init() {
	Config = loadConfiguration()
}