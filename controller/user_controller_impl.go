package controller

import (
	"github/krifik/test-isi/model"
	"github/krifik/test-isi/service"

	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserControllerImpl(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(c *fiber.Ctx) error {
	var request model.CreateUserRequest
	c.BodyParser(&request)

	result, err := controller.UserService.Register(request)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(model.WebResponse{
			Code:   fiber.ErrBadRequest.Code,
			Remark: err.Error(),
			Data:   result,
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(model.WebResponse{
			Code:   fiber.StatusOK,
			Remark: "success",
			Data:   result,
		})
	}
}
