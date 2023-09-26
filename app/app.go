package app

import (
	"github/krifik/test-isi/config"
	"github/krifik/test-isi/controller"
	"github/krifik/test-isi/repository"
	"github/krifik/test-isi/routes"
	"github/krifik/test-isi/service"

	"github.com/gofiber/fiber/v2"
)

func InitializeApp() *fiber.App {
	conf := config.NewConfiguration()
	db := config.NewPostgresDatabase(conf)
	config.NewRunMigration(db)
	// User
	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository)
	userController := controller.NewUserControllerImpl(userService)
	// Account
	accountRepository := repository.NewAccountRepositoryImpl(db)
	accountService := service.NewAccountServiceImpl(accountRepository)
	accountController := controller.NewAccountControllerImpl(accountService)
	// Mutation
	mutationRepository := repository.NewMutationRepositoryImpl(db)
	mutationService := service.NewMutationServiceImpl(mutationRepository)
	mutationController := controller.NewMutationControllerImpl(mutationService)

	app := fiber.New()
	routes.Route(app, userController, accountController, mutationController)

	return app
}
