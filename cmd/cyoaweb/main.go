package main

import (
	"cyoa"
	"flag"
	"fmt"
	"os"
)

func main() {
	parsingJson()
}

func parsingJson() {

	// Create flags for our optional variables
	filename := flag.String("file", "gopher.json", "the json file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	// Open the JSON file and parse ths story in it
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	story, err :=  cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(story)
}