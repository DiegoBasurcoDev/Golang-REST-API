package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
)

type task struct {
	ID int `json:ID`
	Name string `json:Name`
	Content string `json:Content`
}

type allTasks []task

var tasks = allTasks {
	{
		ID: 1,
		Name: "Task 1",
		Content: "Some Content",
	},
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my REST API");
}

func main()  {
	//fmt.Println("Hello World")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute);
	router.HandleFunc("/tasks", getTasks)

	log.Fatal(http.ListenAndServe(":3000", router))
}