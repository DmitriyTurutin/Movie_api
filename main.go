package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "log"
  "net/http"
)

type Movie struct {
  ID string `json:"id"`
  Isbn string `json:"isbn"`
  Title string `json:"title"`
  Director *Director `json:"director"`
}

type Director struct {
  FirstName string `json:"firstname"`
  LastName string `json:"lastname"`
}

var movies []Movie

func main(){
  r := mux.NewRouter()

  movies = append(movies, Movie{ID: "1", Isbn: "1234", Title: "Some movie", Director: &Director{FirstName:"Dmitriy", LastName: "Turutin"}})
  movies = append(movies, Movie{ID: "2", Isbn: "1235", Title: "Some movie about love", Director: &Director{FirstName:"Kate", LastName: "Glusha"}})
  r.HandleFunc("/movies", getMovies).Methods("GET")
  r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
  r.HandleFunc("movies", createMovie).Methods("POST")
  r.HandleFunc("movies/{id}", updateMovie).Methods("PUT")
  r.HandleFunc("movies/{id}", deleteMovie).Methods("DELETE")

  fmt.Println("Starting server at port 7777\n")
  log.Fatal(http.ListenAndServe(":7777",r))

}