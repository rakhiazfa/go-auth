package main

import (
	"fmt"
	"github.com/rakhiazfa/vust-identity-service/internal/config"
	"github.com/rakhiazfa/vust-identity-service/internal/database"
	"github.com/rakhiazfa/vust-identity-service/internal/entities"
	"github.com/rakhiazfa/vust-identity-service/wire"
	"github.com/spf13/viper"
)

func main() {
	config.SetupConfig()

	db := database.NewPostgresConnection()
	r := wire.SetupProviders()

	// Permissions

	var permissions []entities.Permission

	for _, route := range r.Routes() {
		permissions = append(permissions, entities.Permission{
			Name:       fmt.Sprintf("%s %s", route.Method, route.Path),
			ServiceKey: viper.GetString("application.key"),
			Method:     route.Method,
			Path:       route.Path,
		})
	}

	db.Create(permissions)

	// Roles

	roles := []entities.Role{
		{Name: "Super Admin"},
		{Name: "Customer"},
	}

	db.Create(&roles)

	// Users

	users := []entities.User{
		{
			Name:     "Super Admin",
			Username: "super.admin",
			Email:    "super.admin@vust.com",
			Password: "P@ssw0rd",
			Roles:    []entities.Role{roles[0]},
		},
		{
			Name:     "Customer",
			Username: "customer",
			Email:    "customer@vust.com",
			Password: "P@ssw0rd",
			Roles:    []entities.Role{roles[1]},
		},
	}

	db.Create(&users)
}
