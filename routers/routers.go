package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/LibraryApp/handlers"
)

func Router(app *fiber.App) {

	app.Post("/login", handlers.Login)
	app.Post("/sign_up", handlers.SingUp)

	app.Get("/books", handlers.KitapGetAll)
	app.Get("/book/:id", handlers.KitapGetByID)
	app.Post("/book", handlers.KitapAdd)
	app.Delete("/book", handlers.KitapDelete)

	app.Get("/favori/:kullanici_id", handlers.FavoriGetAll)
	app.Post("/favori", handlers.FavoriAdd)
	app.Delete("/favori", handlers.FavoriDelete)

}
