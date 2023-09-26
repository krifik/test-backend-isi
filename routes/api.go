package routes

import (
	"github/krifik/test-isi/controller"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, userController controller.UserController, accountController controller.AccountController, mutationController controller.MutationController) {
	app.Post("/api/daftar", userController.Register)
	app.Post("/api/tabung", accountController.SavingOrWithdraw)
	app.Post("/api/tarik", accountController.SavingOrWithdraw)
	app.Get("/api/saldo/:no_rekening", accountController.FindAccount)
	app.Get("/api/mutasi/:no_rekening", mutationController.FindMutations)
}
