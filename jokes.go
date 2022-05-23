package main

import (
	"fmt"
	"strings"

	"github.com/icelain/jokeapi"
)

type JokesResp struct {
	Error    bool
	Category string
	JokeType string
	Joke     []string
	Flags    map[string]bool
	Id       float64
	Lang     string
}

func Joke() (string, error) {
	jt := "single"
	blacklist := []string{"nsfw"}
	ctgs := []string{"Programming"}

	api := jokeapi.New()

	api.Set(jokeapi.Params{Blacklist: blacklist, JokeType: jt, Categories: ctgs})
	response, err := api.Fetch()
	if err != nil {
		fmt.Println(err)
	}
	return strings.Join(response.Joke, ""), err
}
