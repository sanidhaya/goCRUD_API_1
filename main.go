package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"LastName"`
}

var movies []Movie

func getmovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deletemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, items := range movies {
		if items.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, items := range movies {
		if items.Id == params["id"] {
			json.NewEncoder(w).Encode(items)
			return
		}
	}
}

func createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	for _, items := range movies {
		movie_id := strconv.Itoa(rand.Intn(100000000))
		if items.Id != movie.Id {
			movie.Id = movie_id
			continue
		}
	}
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updatemoive(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, items := range movies {
		if items.Id == params["id"] {
			movies = append(movies[index:], movies[:index+1]...)
			var new_movie Movie
			_ = json.NewDecoder(r.Body).Decode(&new_movie)
			new_movie.Id = params["id"]
			movies = append(movies, new_movie)
			json.NewEncoder(w).Encode(new_movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		Id:    "1",
		Isbn:  "123",
		Title: "movie1",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Doe",
		},
	})

	movies = append(movies, Movie{
		Id:    "2",
		Isbn:  "456",
		Title: "movie2",
		Director: &Director{
			Firstname: "Jane",
			Lastname:  "Smith",
		},
	})

	r.HandleFunc("/movies", getmovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getmovie).Methods("GET")
	r.HandleFunc("/movies", createmovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updatemoive).Methods("PUT")
	r.HandleFunc("/movies/{id}", deletemovie).Methods("DELETE")

	fmt.Println(movies)

	fmt.Printf("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
