package main

import ("fmt"
		"net/http"
		"github.com/gorilla/mux"
		"encoding/json"
		"gopkg.in/mgo.v2"
		//"gopkg.in/mgo.v2/bson"
		)

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if(err != nil){
		panic(err)
	}

	return session
}

var collection = getSession().DB("curso_go").C("movies");

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola mundo desde mi servidor web con go!")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	var results []Movie
	err := collection.Find(nil).Sort("-_id").All(&results)

	if(err != nil){
		w.WriteHeader(500)
		return
	}
	
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	movie_id := params["id"]

	fmt.Fprintf(w, "Cargado la pelicula numero %s",movie_id)	
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if(err != nil){
		panic(err)
	}

	defer r.Body.Close()

	err = collection.Insert(movie_data)

	if(err != nil){
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(movie_data)
}