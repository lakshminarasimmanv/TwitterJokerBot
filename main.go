package main

import "fmt"

func main() {
	joke, err := Joke()
	if err != nil {
		fmt.Println(err)
	} else {
		Twitter(joke)
	}
}
