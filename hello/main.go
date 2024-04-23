package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("how is it goint?")
	r := mux.NewRouter()
	// i can also handle route like this or like below this
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {		//this default which method => GET
		fmt.Fprintf(w, "how is it going hardik")
	})


	// 
	r.HandleFunc("/hello", demo).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))

}

func demo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World HARDIK!"))			//this is the way to write response..passed in byte format. in string format it will not work. we have passed it as a slice of byte
}