package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/LibraryApp/handlers"
)

func Router(app *fiber.App) {

	app.Post("/login", handlers.Login)
	app.Post("/sing-up", handlers.SingUp)

	app.Get("/kitaplar", handlers.KitapGetAll)
	app.Get("/kitap/:id", handlers.KitapGetByID)
	app.Post("/kitap", handlers.KitapAdd)

}
