package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Person struct {
	ID   int
	Name string
	age  int
}

var people []Person

func GetId() int {
	peopleLen := len(people)
	if peopleLen == 0 {
		return 1
	}

	return people[peopleLen-1].ID + 1
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for _, person := range people {
		if strconv.Itoa(person.ID) == id {
			json.NewEncoder(w).Encode(person)
			return
		}
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	var person Person

	_ = json.NewDecoder(r.Body).Decode(&person)

	person.ID = GetId()
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for index, person := range people {
		if string(rune(person.ID)) == id {
			people = append(people[:index], people[index+1:]...)
			return
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()

	for i := 1; i < 5; i++ {
		people = append(people, Person{ID: i, age: 1 * 12, Name: strconv.Itoa(i) + " testes"})
	}

	router.HandleFunc("/contato", GetPeople).Methods("GET")
	router.HandleFunc("/contato/{id}", GetById).Methods("GET")
	router.HandleFunc("/contato", Create).Methods("Post")
	router.HandleFunc("/contato/{id}", Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
