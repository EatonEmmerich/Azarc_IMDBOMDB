package main

import (
    "flag"
    "fmt"

	"Azarc/imdb"
	"Azarc/omdb"
)

var apiKey = flag.String("api_key","","The omdb API key")
var filePath = flag.String("file_path", "title.basics.tsv", "Path to the file containing the imdb names")


func main() {
	flag.Parse()

	imdbClient, err := imdb.New(*filePath)
	list, err := imdbClient.List()
	if err != nil {
		panic(err)
	}
	fmt.Printf("item: %#v", list)

	omdbClient := omdb.New(*apiKey)
	item, err := omdbClient.Info("tt0000005")
	if err != nil {
		panic(err)
	}
	fmt.Printf("item: %#v", item)
}
