package model

import "time"

type Github struct {
	SHA    string `json:"sha"`
	Commit Commit `json:"commit"`
}

type Commit struct {
	Committer Committer `json:"committer"`
	Message   string    `json:"message"`
}

type Committer struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

type Data struct {
	App string
	Committer string
	Message   string
	SHA       string
	Date      time.Time
}