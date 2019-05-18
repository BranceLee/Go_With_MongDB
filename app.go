package main 

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"github.com/Plee/Mongolang/models"
	"github.com/Plee/Mongolang/api"
	"github.com/Plee/Mongolang/config"
)

// 实例化后端MovieDao
var movieApi =api.MoviesAPI{}
var configer = config.Config{}

func AllMovies(w http.ResponseWriter,r *http.Request){
	// fmt.Fprintln(w, "not implemented yet !")
	defer r.Body.Close()
	movies, err:=movieApi.FindAll()
	if err !=nil {
		respondWithError(w, http.StatusBadRequest, "Can not find the data")
		return
	}
	respondWithJson(w, http.StatusOK,movies)
}

func FindMovie(w http.ResponseWriter,r *http.Request){
	params := mux.Vars(r)
	movie, err :=movieApi.FindById(params["id"])
	fmt.Printf("%s\n",movie)
	if err != nil{
		respondWithError(w, http.StatusBadRequest,"Invalid Movie Id ")
		return
	}
	respondWithJson(w, http.StatusOK,movie)
}

func CreateMovie(w http.ResponseWriter,r *http.Request){
	defer r.Body.Close()
	var movie models.Movies
	err := json.NewDecoder(r.Body).Decode(&movie)
	fmt.Printf("%s\n",movie)
	if err != nil {
		respondWithError(w, http.StatusBadRequest,"Invalid request payload")
	}
	movie.ID = bson.NewObjectId()
	if err := movieApi.Create(movie); err !=nil{
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, movie)
}

func respondWithError(w http.ResponseWriter,code int, msg string){
	respondWithJson(w,code,map[string]string{"error":msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}){
	res,_ :=json.Marshal(payload)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(res)
}

// 解析配置文件"config.toml", 初始化建立数据库连接
func init(){
	configer.Read()

	movieApi.Server = configer.Server
	movieApi.Database = configer.Database
	movieApi.Connect()
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/movies",AllMovies).Methods("GET")
	r.HandleFunc("/movies/new",CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",FindMovie).Methods("GET")
	if err:=http.ListenAndServe(":8888",r); err!=nil{
		log.Fatal(err)
	}
}