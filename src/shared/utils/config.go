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
	AccessTokenDuration       time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration      time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()

	// tự động ghi đè các giá trị mà nó đã đọc từ tệp cấu hình bằng các giá trị của các biến môi trường tương ứng nếu chúng tồn tại.
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	//Unmarshal sắp xếp lại các giá trị vào đối tượng đích config
	err = viper.Unmarshal(&config)
	return
}
