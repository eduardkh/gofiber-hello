package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// app.Static("/", "./public")
	// => http://localhost:80/js/script.js
	// => http://localhost:80/css/style.css

	// app.Static("/prefix", "./public")
	// => http://localhost:80/prefix/js/script.js
	// => http://localhost:80/prefix/css/style.css

	app.Static("*", "./public/index.html")
	// => http://localhost:80/any/path/shows/index/html

	log.Fatal(app.Listen(":80"))
}
