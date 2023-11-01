package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in Golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hello from the Greeter function!")

}
func serveHome(w http.ResponseWriter, r *http.Request) { //designed for gorilla mux
	w.Write([]byte("<h1>Welcome to Golang Series on YT.</h1>"))
}
