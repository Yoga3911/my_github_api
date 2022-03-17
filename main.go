package main

import (
	"github_api/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)



func main() {
	app := fiber.New()

	routes.Data(app)

	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
