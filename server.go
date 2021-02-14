package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleRequest() {
	fs := http.FileServer(http.Dir("./web/ang-client/dist/my-app/"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("/articles", returnAllArticles)
}

func main() {
	Articles = []Article{
		{PokemonName: "Pikachu"},
		{PokemonName: "Charmander"},
	}
	handleRequest()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Article struct {
	PokemonName string `json:"Pokemon"`
}

var Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: retturnAllArtticles")
	json.NewEncoder(w).Encode(Articles)
}
