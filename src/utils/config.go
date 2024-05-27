package utils

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver                  string        `mapstructure:"DB_DRIVER"`
	DBSource                  string        `mapstructure:"DB_SOURCE"`
	HTTPServerAddress         string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress         string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	HTTPServerAddressCustomer string        `mapstructure:"HTTP_SERVER_ADDRESS_CUSTOMER"`
	GRPCServerAddressCustomer string        `mapstructure:"GRPC_SERVER_ADDRESS_CUSTOMER"`
	TokenSymmetricKey         string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	TokenSymmetricKeyCustomer string        `mapstructure:"TOKEN_SYMMETRIC_KEY_CUSTOMER"`
	AccessTokenDuration       time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration      time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	AesKey                    string        `mapstructure:"AES_KEY"`
	AccountEmail              string        `mapstructure:"ACCOUNT_EMAIL"`
	PasswordEmail             string        `mapstructure:"PASSWORD_EMAIL"`
}

func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
