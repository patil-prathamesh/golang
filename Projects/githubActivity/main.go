package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type GithubEvent struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Actor struct {
		Id string `json:"id"`
		Login string `json:"login"`
		DisplayLogin string `json:"display_login"`
		URL string `json:"url"`
		AvatarUrl string `json:"avatar_url"`
	} `json:"actor"`
	Repo struct {
		Id string `json:"id"`
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"repo"`
	Payload struct {
		Head string `json:"head"`
		Before string `json:"before"`
		Commits []struct{
			Sha string `json:"sha"`
		} `json:"commits"`
	} `json:"payload"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No user")
	}

	url := fmt.Sprintf("https://api.github.com/users/%v/events",os.Args[1])
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Err0r")
	}

	var events []GithubEvent
	body, err := io.ReadAll(res.Body)
	if err := json.Unmarshal(body, &events); err != nil {
		fmt.Println("error")
	}

	for i, v := range events{
		if i == 4 {
			return
		}
		fmt.Println(i," -> ",v.Repo.Name)
	}

}