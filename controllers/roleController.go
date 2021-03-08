package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/grak0s/back-golang-postgress/database"
	"github.com/grak0s/back-golang-postgress/models"
)

//AllRoles trase todos los usuarios de la BD
func AllRoles(c *fiber.Ctx) error {

	var role []models.Role

	database.DB.Find(&role)

	return c.JSON(role)
}

//CreateRole crea usuarios
func CreateRole(c *fiber.Ctx) error {
	var roleDTO fiber.Map

	if err := c.BodyParser(&roleDTO); err != nil {
		return err
	}

	//var permissions []models.Permission

	list := roleDTO["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name:       roleDTO["name"].(string),
		Permission: permissions,
	}

	database.DB.Create(&role)

	return c.JSON(role)
}

//GetRole busca un usuario
func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Preload("permissions").Find(&role)

	return c.JSON(role)

}

//UpdateRole actualiza al usuario
func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var roleDTO fiber.Map

	if err := c.BodyParser(&roleDTO); err != nil {
		return err
	}

	//var permissions []models.Permission

	list := roleDTO["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	var result interface{}

	database.DB.Table("role_permissions").Where("role_id", id).Delete(&result)

	role := models.Role{
		Id:         uint(id),
		Name:       roleDTO["name"].(string),
		Permission: permissions,
	}

	database.DB.Model(&role).Updates(role)

	return c.JSON(role)

}

//DeleteRole elimina el usuario
func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Delete(&role)
	return nil
}
