package apps

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"github_api/model"
)

func Indogram(c *fiber.Ctx) error {
	_, resp, err := fasthttp.Get(nil, "https://api.github.com/repos/Yoga3911/indogram/commits/main")
	if err != nil {
		log.Println(err)
	}

	var github model.Github
	json.Unmarshal(resp, &github)

	var data model.Data
	data.App = "Indogram"
	data.Committer = github.Commit.Committer.Name
	data.Message = github.Commit.Message
	data.SHA = github.SHA
	data.Date = github.Commit.Committer.Date

	tmp := fmt.Sprintf("App Name: %v\nCommitter: %v\nCommit: %v\nDate: %v\nSHA: %v", data.App, data.Committer, data.Message, data.Date, data.SHA)

	return c.Status(fiber.StatusOK).SendString(tmp)
}

func Breadify(c *fiber.Ctx) error {
	_, resp, err := fasthttp.Get(nil, "https://api.github.com/repos/Yoga3911/pbm-a1/commits/master")
	if err != nil {
		log.Println(err)
	}

	var github model.Github
	json.Unmarshal(resp, &github)

	var data model.Data
	data.App = "Breadify"
	data.Committer = github.Commit.Committer.Name
	data.Message = github.Commit.Message
	data.SHA = github.SHA
	data.Date = github.Commit.Committer.Date

	tmp := fmt.Sprintf("App Name: %v\nCommitter: %v\nCommit: %v\nDate: %v\nSHA: %v", data.App, data.Committer, data.Message, data.Date, data.SHA)

	return c.Status(fiber.StatusOK).SendString(tmp)
}

func Sulai(c *fiber.Ctx) error {
	_, resp, err := fasthttp.Get(nil, "https://api.github.com/repos/Yoga3911/sulai/commits/master")
	if err != nil {
		log.Println(err)
	}

	var github model.Github
	json.Unmarshal(resp, &github)

	var data model.Data
	data.App = "Sulai"
	data.Committer = github.Commit.Committer.Name
	data.Message = github.Commit.Message
	data.SHA = github.SHA
	data.Date = github.Commit.Committer.Date

	tmp := fmt.Sprintf("App Name: %v\nCommitter: %v\nCommit: %v\nDate: %v\nSHA: %v", data.App, data.Committer, data.Message, data.Date, data.SHA)

	return c.Status(fiber.StatusOK).SendString(tmp)
}