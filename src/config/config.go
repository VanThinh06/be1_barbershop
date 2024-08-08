package config

import (
	"context"
	"fmt"
	"os"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
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
	CloudinaryUrl             string        `mapstructure:"CLOUDINARY_URL"`
	CloudinaryName            string        `mapstructure:"CLOUDINARY_NAME"`
	CloudinarySecret          string        `mapstructure:"CLOUDINARY_SECRET_KEY"`
	CloudinaryKey             string        `mapstructure:"CLOUDINARY_KEY"`
}

func InitConfig(path string) (config *Config, err error) {

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
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return
}

func InitLogger() *logrus.Logger {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.SetOutput(os.Stdout)

	logger.SetLevel(logrus.InfoLevel)
	return logger
}

func InitFirebase(ctx context.Context) (*firebase.App, error) {
	firebase, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, err
	}
	return firebase, nil
}

func InitCloudinary(config Config) (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromParams(config.CloudinaryName, config.CloudinaryKey, config.CloudinarySecret)
	if err != nil {
		return nil, err
	}
	return cld, nil
}
