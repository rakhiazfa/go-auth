package main

import (
	"fmt"
	"github.com/rakhiazfa/vust-identity-service/internal/config"
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
	"github.com/rakhiazfa/vust-identity-service/wire"
	"github.com/spf13/viper"
)

func main() {
	config.SetupConfig()
	utils.CatchError(
		wire.SetupProviders().Run(
			fmt.Sprintf("%s:%d", viper.GetString("application.host"), viper.GetInt("application.port")),
		),
	)
}
