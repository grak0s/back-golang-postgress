package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grak0s/back-golang-postgress/utils"
)

//IsAuthenticated verifica que el token jwt sea valido
func IsAuthenticated(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	if _, err := utils.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"menssage": "unauthenticated",
		})
	}

	return c.Next()

}
