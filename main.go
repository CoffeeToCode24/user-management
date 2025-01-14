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

	var users = []User{{"John", 20, 1}, {"Doe", 30, 2}, {"Smith", 40, 3}, {"Tom", 50, 4}}

	app.Get("/user", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})

	app.Post("/user/create", func(c *fiber.Ctx) error {
		var newUser User

		if err := c.BodyParser(&newUser); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
		}

		for _, user := range users {
			if user.ID == newUser.ID {
				return c.Status(400).JSON(fiber.Map{"error": "ID must be unique"})
			}
		}

		users = append(users, newUser)

		return c.Status(200).JSON(users)

	})

	app.Get("/user/getAll", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})

	app.Get("/user/get/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
		}

		for _, user := range users {
			if user.ID == id {
				return c.JSON(user)
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "User not found"})

	})

	app.Post("/user/update/:id" 


		}



	app.Listen(":3100")
}
