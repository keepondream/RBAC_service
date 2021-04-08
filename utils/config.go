package utils

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	REDIS_HOST     string `mapstructure:"REDIS_HOST"`
	REDIS_PORT     string `mapstructure:"REDIS_PORT"`
	REDIS_DB       string `mapstructure:"REDIS_DB"`
	REDIS_PASSWORD string `mapstructure:"REDIS_PASSWORD"`
	REDIS_PREFIX   string `mapstructure:"REDIS_PREFIX"`

	DB_USER string `mapstructure:"DB_USER"`
	DB_PWD  string `mapstructure:"DB_PWD"`
	DB_HOST string `mapstructure:"DB_HOST"`
	DB_PORT string `mapstructure:"DB_PORT"`
	DB_NAME string `mapstructure:"DB_NAME"`

	LOG_LEVEL           string `mapstructure:"LOG_LEVEL"`
	LOG_FILE            string `mapstructure:"LOG_FILE"`
	ENFORCER_MODEL_FILE string `mapstructure:"ENFORCER_MODEL_FILE"`
	MIGRATIONS_TABLE    string `mapstructure:"MIGRATIONS_TABLE"`
	SERVER_PORT         string `mapstructure:"SERVER_PORT"`
	TIME_OUT            string `mapstructure:"TIME_OUT"`
}

// LoadConfig configFile ./app.env or ./app.yaml
// can use JSON, TOML, YAML, HCL, envfile and Java properties config files
func LoadConfig(configFile string) (config *Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(configFile)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(&config); err != nil {
			log.Fatal("configf file changed failed err is ", err)
		}
	})

	return
}
