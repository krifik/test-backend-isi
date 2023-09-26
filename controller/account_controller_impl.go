package controller

import (
	"github/krifik/test-isi/model"
	"github/krifik/test-isi/service"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AccountControllerImpl struct {
	AccountService service.AccountService
}

func NewAccountControllerImpl(accountService service.AccountService) AccountController {
	return &AccountControllerImpl{
		AccountService: accountService,
	}
}

// SavingOrWithdraw implements AccountController
func (controller *AccountControllerImpl) SavingOrWithdraw(c *fiber.Ctx) error {
	var req model.AccountRequest
	err := c.BodyParser(&req)
	var isSaving bool
	if strings.Contains("/api/tabung", c.Path()) {
		isSaving = true
	} else {
		isSaving = false
	}
	result, err := controller.AccountService.SavingOrWithdraw(req, isSaving)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:   fiber.StatusBadRequest,
			Remark: err.Error(),
			Data:   nil,
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(model.WebResponse{
			Code:   fiber.StatusOK,
			Remark: "Success Saving Or Withdraw",
			Data:   result,
		})
	}
}

func (controller *AccountControllerImpl) FindAccount(c *fiber.Ctx) error {
	var req model.AccountRequest
	req.AccountNumber = c.Params("no_rekening")
	result := controller.AccountService.FindAccount(req)
	if result == nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:   fiber.StatusBadRequest,
			Remark: "Account dosent exist",
			Data:   nil,
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(model.WebResponse{
			Code:   fiber.StatusOK,
			Remark: "Success Get Account",
			Data:   result,
		})
	}
}
