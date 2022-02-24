package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/LibraryApp/handlers"
)

func Router(app *fiber.App) {

	app.Get("/login", handlers.Login)

}
