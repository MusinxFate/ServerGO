package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	PokemonName string `json:"Pokemon"`
}

var Articles []Article

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: retturnAllArtticles")
	json.NewEncoder(w).Encode(Articles)
}
