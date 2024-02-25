package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	parsingJson()
}

func parsingJson() {

	port := flag.Int("porn", 3000, "the port to start the CYOA web app on")

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
	
	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on the port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}