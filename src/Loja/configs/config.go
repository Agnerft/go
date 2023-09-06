package configs

import (
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIconfig
	DB  DBconfig
}

type APIconfig struct {
	Port string
}

type DBconfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func Init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.host", "localhost")

}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIconfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBconfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.database"),
	}

	return nil
}
func GetDB() DBconfig {
	return cfg.DB

}

func GetServerPort() string {
	return cfg.API.Port
}
