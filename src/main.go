package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	// app.Use(cors.New())
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "*",
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// }))

	app.Static("/", "./public")
	log.Fatal(app.Listen(":80"))
}
