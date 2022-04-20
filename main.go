package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/omerfruk/LibraryApp/database"
	"github.com/omerfruk/LibraryApp/routers"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Content-Length, Accept-Encoding",
		MaxAge:           86400,
		AllowMethods:     "POST, GET, OPTIONS, PUT, DELETE, UPDATE",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}))

	database.ConnectDB()
	fmt.Println("DB connected")
	database.AutoMigrate()
	database.DefaultUserCreate()

	routers.Router(app)
	port := "3000"

	fmt.Println("Backend Started")
	app.Listen(fmt.Sprintf(":%s", port))

}
