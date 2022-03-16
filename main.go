package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
)

type Github struct {
	SHA    string `json:"sha"`
	Commit Commit `json:"commit"`
}

type Commit struct {
	Committer Committer `json:"committer"`
	Message   string    `json:"message"`
}

type Committer struct {
	Name    string    `json:"name"`
	Date    time.Time `json:"date"`
}

type Data struct {
	SHA       string
	Message   string
	Committer string
	Date      time.Time
}

func getData(c *fiber.Ctx) error {
	_, resp, err := fasthttp.Get(nil, "https://api.github.com/repos/Yoga3911/indogram/commits/main")
	if err != nil {
		log.Println(err)
	}

	var github Github
	json.Unmarshal(resp, &github)

	var data Data
	data.SHA = github.SHA
	data.Committer = github.Commit.Committer.Name
	data.Message = github.Commit.Message
	data.Date = github.Commit.Committer.Date

	tmp := fmt.Sprintf("%v\n%v\n%v\n%v", data.SHA, data.Committer, data.Message, data.Date)

	return c.Status(fiber.StatusOK).SendString(tmp)
}

func main() {
	app := fiber.New()

	app.Get("/api/yoga3911", getData)

	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
