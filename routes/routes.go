package routes

import "github.com/gofiber/fiber/v2"

func Handlers(){
	app := fiber.New()
	app.Get("/user",GetAllUsers)
	app.Listen(":3100")
}