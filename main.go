package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id 				string `json:"Id"`
	Title 		string `json:"Title"`
	Desc 			string `json:"desc"`
	Content 	string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			/*
			 This function encodes `article` as JSON and writes it directly to the HTTP response
			 which is why no explicit return is needed.
			*/
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	log.Fatal(http.ListenAndServe(":10000", myRouter))

	// http.HandleFunc("/", homePage)
	// http.HandleFunc("/articles", returnAllArticles)
	// log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	fmt.Println("Go server running...")

	Articles = []Article{
		{ Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content" },
		{ Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content" },
	}

	handleRequests()
}
