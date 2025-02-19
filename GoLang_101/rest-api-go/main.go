package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	//Id       int64  `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Age      int    `json:"age"`
}

var (
	Pessoa1 = User{Name: "Arthur", Lastname: "Felipe", Age: 20}
)

func main() {
	listUsers := []User{Pessoa1}
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/user", func(user *fiber.Ctx) error {
		return user.JSON(listUsers)
	})
	app.Post("/user", func(user *fiber.Ctx) error {
		return user.JSON(listUsers)
	})
	app.Patch("/user", func(user *fiber.Ctx) error {
		return user.JSON(listUsers)
	})
	app.Delete("/user", func(user *fiber.Ctx) error {
		return user.JSON(listUsers)
	})

	log.Fatal(app.Listen(":3000"))
}
