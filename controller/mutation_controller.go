package controller

import "github.com/gofiber/fiber/v2"

type MutationController interface {
	FindMutations(c *fiber.Ctx) error
}
