package main

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name string `json:"name"`
    Age  int    `json:"age"`
    ID   int    `json:"id"`
}
 

func main() {
	
	app := fiber.New()

	var users = []User{{"John", 20, 1}, {"Doe", 30, 2}, {"Smith", 40, 3}}

	app.Get("/", func(c *fiber.Ctx) error {
		 return  c.JSON(users)
	})
	
	app.Listen(":3100")
}
