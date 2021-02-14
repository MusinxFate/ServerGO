package main

import (
	"log"
	"net/http"
)

func handleRequest() {
	fs := http.FileServer(http.Dir("./web/ang-client/dist/my-app/"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("/articles", ReturnAllArticles)
}

func main() {
	Articles = []Article{
		{PokemonName: "Pikachu"},
		{PokemonName: "Charmander"},
	}
	handleRequest()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
