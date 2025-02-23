package main

import (
	"log"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Age      int    `json:"age"`
}

var (
	Pessoa1 = User{Name: "Arthur", Lastname: "Felipe", Age: 20}
	Pessoa2 = User{Name: "Amanda", Lastname: "Silva", Age: 20}
	Pessoa3 = User{Name: "Felipe", Lastname: "Silva", Age: 20}
	Pessoa4 = User{Name: "Jurandir", Lastname: "Silva", Age: 20}
)

var mySwaggerByteSlice = []byte(`{
    "swagger": "2.0",
    "info": {
        "title": "Swagger API Docs",
        "description": "API documentation",
        "version": "1.0.0"
    },
    "paths": {}
}`)

func main() {
	listUsers := []User{Pessoa1, Pessoa2, Pessoa3, Pessoa4}
	app := fiber.New()

	app.Use("/swagger", swagger.New(swagger.Config{
		BasePath:    "/",
		FilePath:    "./docs/swagger.json",
		FileContent: mySwaggerByteSlice,
		Path:        "swagger",
		Title:       "Swagger API Docs",
	}))

	app.Post("/swagger/user", func(c *fiber.Ctx) error {
		return c.JSON(listUsers)
	})
	app.Patch("/swagger/user", func(c *fiber.Ctx) error {
		return c.JSON(listUsers)
	})
	app.Delete("/swagger/user", func(c *fiber.Ctx) error {
		return c.JSON(listUsers)
	})

	log.Fatal(app.Listen(":3000"))
}
