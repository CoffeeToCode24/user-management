package main

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	ID   int    `json:"id"`
}

var users = []User{{"John", 20, 1}, {"Doe", 30, 2}, {"Smith", 40, 3}, {"Tom", 50, 4}}

//	app.Get("/user", func(c *fiber.Ctx) error {
//		?return c.JSON(users)
//		})
func getAllUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}

// app.Post("/user/create", func(c *fiber.Ctx) error {
func createUser(c *fiber.Ctx) error {
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
}

//	app.Get("/user/getAll", func(c *fiber.Ctx) error {
//		return c.JSON(users)
//	})
func getAll(c *fiber.Ctx) error {
	return c.JSON(users)
}

// app.Get("/user/get/:id", func(c *fiber.Ctx) error {
func getUserByID(c *fiber.Ctx) error {
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
}

func updateUserByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var updatedUser User

	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Input"})
	}

	for i, user := range users {
		if user.ID == id {
			users[i].Name = updatedUser.Name
			users[i].Age = updatedUser.Age
			return c.JSON(fiber.Map{"success": "user updated successfully", "user": users[i]})
		}

	}
	return c.Status(400).JSON(fiber.Map{"error": "no user found"})
}

func deleteUserByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)

			return c.JSON(fiber.Map{"success": "user deleted successfully"})
		}

	}
	return c.Status(400).JSON(fiber.Map{"error": "no user found"})
}

func main() {
	app := fiber.New()

	app.Get("/user", getAllUsers)

	app.Post("/user/create", createUser)

	app.Get("/user/getAll", getAllUsers)

	app.Get("/user/get/:id", getUserByID)

	app.Post("/user/update/:id", updateUserByID)

	app.Delete("/user/delete/:id", deleteUserByID)

	app.Listen(":3100")
}
