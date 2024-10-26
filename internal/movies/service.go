package movies

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

var movies []Movie = []Movie{
	{"1", "43277", "Batman", &Director{FirstName: "Cristoph", LastName: "Nolan"}},
	{"2", "43278", "Batman", &Director{FirstName: "Cristoph", LastName: "Nolan"}},
	{"2", "43279", "Batman", &Director{FirstName: "Cristoph", LastName: "Nolan"}},
	{"3", "43280", "Batman", &Director{FirstName: "Cristoph", LastName: "Nolan"}},
}

func GetMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				return
			}
			return
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000000) + 1)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(&movie)
		}
	}
}
