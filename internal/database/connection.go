package database

import (
	"fmt"
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresConnection() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Jakarta",
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.name"),
		viper.GetString("database.username"),
		viper.GetString("database.password"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	utils.CatchError(err)

	return db
}
