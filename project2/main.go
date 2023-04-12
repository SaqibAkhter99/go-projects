package main 

import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct { 
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title`
	Director *Director `json:director`
}
type Director struct {
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
}

var movies[]Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	param := mux.Vars(r)
	for id,value := range movies{
		if value.ID == param["id"]{
			movies = append(movies[:id], movies[id+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(movies)
}

func createMovies(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-type","application/json")
	var movie Movie 
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000000))
	movies = append(movies,movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	param:=mux.Vars(r)
	for id,value := range movies{
		if value.ID == param["id"]{
			movies = append(movies[:id], movies[id+1:]...)
			var movie Movie 
			_ = json.NewDecoder(r.Body).Decode(&movie)
	        movie.ID = param["id"]
	        movies = append(movies,movie)
	        json.NewEncoder(w).Encode(movies)
			break
		}

	}

}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	param := mux.Vars(r)
	for _,value:= range movies{
		if value.ID == param["id"]{
			json.NewEncoder(w).Encode(value)
			return 
		}
	}
}

func main(){
	r:=mux.NewRouter()
	movies = append(movies, Movie{
		ID      :"1",
		Isbn    :"438277",
		Title   :"The GodFather",
		Director:&Director {
			First_Name:"Saqib",
			Last_Name : "Akhter",
		}, 
	})
	movies = append(movies, Movie{
		ID      :"2",
		Isbn    :"438288",
		Title   :"Hera Pheri",
		Director:&Director{
			First_Name:"Farzan",
			Last_Name : "Alnoor",
		}, 
	})
	r.HandleFunc("/movies"     ,   getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",    getMovie).Methods("GET")
	r.HandleFunc("/movies"     ,createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Println("Starting the Server at Port 8000")
	log.Fatal(http.ListenAndServe(":8000",r))

}