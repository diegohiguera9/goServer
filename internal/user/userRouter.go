package user

import (
	"github.com/gofiber/fiber/v2"
)

var UserGroup fiber.Router

func UserRouter() {

	UserGroup.Get("/create", CreateUser)

}
