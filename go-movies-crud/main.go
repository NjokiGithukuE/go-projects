package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.Response, r*http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.Response, r*http.Request) {
	w.Header().set("Content-Type", "application/json")
	params := 
}

func getMovie(w http.Response, r*http.Request) {

}

func createMovie(w http.Response, r*http.Request) {

}

func updateMovie(w http.Response, r*http.Request) {
	
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie[ID: "1", Isbn: "438227", Title: "Movie One", Director: %Director[Firstname: "John", Lastname: "Doe"]])
	movies = append(movies, Movie[ID: "2", Isbn: "45455", Title: "Movie Two", Director: %Director[Firstname: "Steve", Lastname: "Smith"]])
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/(id)", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/(id)", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/(id)", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}