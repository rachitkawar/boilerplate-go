package common

import (
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		Log.Fatalf("Error reading config file: %v", err)
	}
}

func GetEnv(key string) string {
	return viper.GetString(key)
}
