package config

import (
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
	"github.com/spf13/viper"
)

func SetupConfig() {
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")

	utils.CatchError(viper.ReadInConfig())
}
