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
	port := flag.Int("port", 3000, "the port to start CYOA web application on")
	file := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	flag.Parse()

	fmt.Printf("Using the sotry in %s.\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		log.Fatal("error opening file:", *file)
	}
	defer f.Close()

	story, err := cyoa.JsonStory(f)
	if err != nil {
		log.Fatal("Unable to parse to JSON.\nerr:", err)
	}

	h := cyoa.NewHandler(story)
	fmt.Println("Starting server on port: ", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
