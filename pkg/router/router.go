package router

import (
	"github.com/gorilla/mux"
	"go.mod/internal/movies"
)

func SetupeRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/movies", movies.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", movies.GetMovie).Methods("GET")
	r.HandleFunc("/movies", movies.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", movies.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", movies.DeleteMovie).Methods("DELETE")

	return r
}
