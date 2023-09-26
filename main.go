package main

import (
	"github/krifik/test-isi/app"
	"os"
)

func main() {
	app := app.InitializeApp()

	app.Listen(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT"))
}
