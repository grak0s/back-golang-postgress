package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grak0s/back-golang-postgress/database"
	"github.com/grak0s/back-golang-postgress/models"
)

//AllPermission trase todos los permisos  de la BD
func AllPermission(c *fiber.Ctx) error {

	var permission []models.Permission

	database.DB.Find(&permission)

	return c.JSON(permission)
}
