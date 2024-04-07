package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
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
	dst, err := os.Create("../fileupload_go/" + handler.Filename)
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
	w.Write([]byte("File uploaded successfully"))
}

func main() {
	// Define the route for the file upload endpoint
	http.HandleFunc("/upload", uploadFile)

	// Start the HTTP server on port 8080
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
