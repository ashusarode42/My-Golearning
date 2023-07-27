package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Print("Hello mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveName).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))

}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series</h1>"))
}
