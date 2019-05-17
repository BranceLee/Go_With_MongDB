package main 

import (
	// "fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func AllMovies(w http.ResponseWriter,r *http.Request){
	return
}

func FindMovie(w http.ResponseWriter,r *http.Request){
	return
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/movies",AllMovies).Methods("GET")
	r.HandleFunc("movies/{id}",FindMovie).Methods("GET")
	if err:=http.ListenAndServe(":5000",nil); err!=nil{
		log.Fatal(err)
	}
}