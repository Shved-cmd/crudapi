package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// CrudAPI

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438228", Title: "Movie Two", Director: &Director{FirstName: "Michael", LastName: "Jackson"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies{id}", deleteMovie).Methods("DELETE")

	log.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {

}

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"` // уникальный идентификатор фильма
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie
