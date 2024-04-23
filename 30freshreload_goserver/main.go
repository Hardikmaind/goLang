package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"log"
	
	
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received uploadFile API call: %s %s", r.Method, r.URL.Path)
	// Parse the multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		return
	}

	// Retrieve the file from the form data
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file from form data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a new file on the server
	dst, err := os.Create("../30freshreload_goserver/" + handler.Filename)
	if err != nil {
		http.Error(w, "Unable to create file on server", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the file data to the destination file on the server
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error copying file data to server", http.StatusInternalServerError)
		return
	}

	// Send a success response to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded successfully Hardik"))
}

func main() {
	// Define the route for the file upload endpoint
	starttime := time.Now()
	http.HandleFunc("/upload", uploadFile)
	endtime := time.Now()

	fmt.Println("Time taken to start the server is ", endtime.Sub(starttime))

	// Start the HTTP server on port 8080
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)






	// install the air with the help of go intall {air repo}
	//then do air init ...this will create the .air.toml file
	// then to start the server simple press air=========>>>>server will simply start with auto reloading
}



// steps to start server with the fresh
// first install the fresh with => go install github.com/zzwx/fresh@latest
// then do fresh -g to generate the fresh.yaml file
// then start the server witht the fresh keyword