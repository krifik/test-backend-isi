package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	Register(c *fiber.Ctx) error
}
