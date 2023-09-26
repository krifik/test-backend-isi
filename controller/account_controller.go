package controller

import "github.com/gofiber/fiber/v2"

type AccountController interface {
	SavingOrWithdraw(c *fiber.Ctx) error
	FindAccount(c *fiber.Ctx) error
}
