package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to LearnCodeonline server"))
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{"message": "Hello from learnCodeonline.in"}
		jsonResponse, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	})

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		var myJson map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&myJson)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		jsonResponse, _ := json.Marshal(myJson)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	})

	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		formData := make(map[string]interface{})
		for key, values := range r.PostForm {
			formData[key] = values[0]
		}
		jsonResponse, _ := json.Marshal(formData)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	})

	port := 8000
	fmt.Printf("Example app listening at http://localhost:%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}
