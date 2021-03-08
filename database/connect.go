package database

import (
	"github.com/grak0s/back-golang-postgress/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

//Connect conexión a la base de datos de postgres
func Connect() {

	dsn := "host=localhost user=postgres password=postgres dbname=go_admin port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("no se conecta a la base de datos")
	}

	DB = database

	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})

}
