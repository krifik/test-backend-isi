package controller

import (
	"github/krifik/test-isi/model"
	"github/krifik/test-isi/service"

	"github.com/gofiber/fiber/v2"
)

type MutationControllerImpl struct {
	MutationService service.MutationService
}

// FindMutations implements MutationController
func (controller *MutationControllerImpl) FindMutations(c *fiber.Ctx) error {
	var req model.MutationRequest
	req.AccountNumber = c.Params("no_rekening")
	result := controller.MutationService.FindMutations(req)
	if result != nil {
		return c.Status(fiber.StatusOK).JSON(model.WebResponse{
			Code:   fiber.StatusOK,
			Remark: "OK",
			Data:   result,
		})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:   fiber.StatusBadRequest,
			Remark: "Not Found",
			Data:   nil,
		})
	}
}

func NewMutationControllerImpl(mutationService service.MutationService) MutationController {
	return &MutationControllerImpl{
		MutationService: mutationService,
	}
}
