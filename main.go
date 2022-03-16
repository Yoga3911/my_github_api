package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Github struct {
	SHA string `json:"sha"`
	Commit Commit `json:"commit"`
}

type Commit struct {
	Committer Committer `json:"committer"` 
	Message string `json:"message"`
}

type Committer struct {
	Name string `json:"name"`
	Date time.Time `json:"date"`
	NewDate string 
}

func main() {
	resp, err := http.Get("https://api.github.com/repos/Yoga3911/indogram/commits/main")
	if err != nil {
		log.Println(err);
	}

 	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var data Github
	json.Unmarshal(respData, &data)

	data.Commit.Committer.NewDate = data.Commit.Committer.Date.Format("02-01-2006")

	fmt.Println(data.Commit.Committer.NewDate)
}