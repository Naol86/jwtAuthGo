package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerPort             string `mapstructure:"SERVER_PORT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPassword             string `mapstructure:"DB_PASSWORD"`
	DBName                 string `mapstructure:"DB_NAME"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	AccessTokenExpireHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRE_HOUR"`
	RefreshTokenExpireHour int    `mapstructure:"REFRESH_TOKEN_EXPIRE_HOUR"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
}

func NewEnv() (*Env, error) {
	env := Env{}
	v := viper.New()

	// Tell Viper we’re using a .env file
	v.SetConfigFile(".env")
	v.SetConfigType("env")      // important for .env files
	v.AddConfigPath(".")        // look in current directory
	v.AddConfigPath("./config") // (optional) if your .env is in /config
	v.AutomaticEnv()            // read system ENV vars too

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	if err := v.Unmarshal(&env); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	if env.AppEnv == "development" {
		log.Println("Running in development mode")
	}

	return &env, nil
}
