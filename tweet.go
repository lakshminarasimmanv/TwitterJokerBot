package main

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func Twitter(j string) {

	config := oauth1.NewConfig(os.Getenv("CONKEY"), os.Getenv("CONSEC"))
	token := oauth1.NewToken(os.Getenv("ACTO"), os.Getenv("ACTOSEC"))

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	_, _, err := client.Statuses.Update(j, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
