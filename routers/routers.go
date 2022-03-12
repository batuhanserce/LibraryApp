package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/LibraryApp/handlers"
)

func Router(app *fiber.App) {

	app.Post("/login", handlers.Login)
	app.Post("/sign_up", handlers.SingUp)

	app.Get("/book", handlers.KitapGetAll)
	app.Get("/book/:id", handlers.KitapGetByID)
	app.Post("/book", handlers.KitapAdd)
	app.Delete("/book", handlers.KitapDelete)

	app.Get("/katagori", handlers.KatagoriGetAll)
	app.Get("/katagori/:id", handlers.KatagoriGetByID)
	app.Post("/katagori", handlers.KatagoriAdd)
	app.Delete("/katagori", handlers.KatagoriDelete)

	app.Get("/user", handlers.KullaniciGetAll)
	app.Get("/user/:id", handlers.KullaniciGetByID)
	app.Post("/user", handlers.KullaniciAdd)
	app.Delete("/user", handlers.KullaniciDelete)

	app.Get("/favori", handlers.FavoriGetAll)
	app.Get("/favori/:kullanici_id", handlers.FavoriGetAllByKullaniciId)
	app.Post("/favori", handlers.FavoriAdd)
	app.Delete("/favori", handlers.FavoriDelete)

}
