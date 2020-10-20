package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article abc
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

// Articles abc
type Articles []Article

func main() {
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{Article{Title: "Title 2", Desc: "Desc 2", Content: "Content 2"}}

	fmt.Println("Hit: all articles endpoint")
	json.NewEncoder(w).Encode(articles)
}
