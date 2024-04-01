package main

import (
	"fmt"
	// "io/ioutil"			//this is deprecreated, instead use io.ReadAll for reading the data
	"net/http"
	"io"
)

const url = "https://official-joke-api.appspot.com/random_joke"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() // caller's responsibility to close the connection

	// databytes, err := ioutil.ReadAll(response.Body)			//this is deprecreated, instead use io.ReadAll
	databytes2,err2:= io.ReadAll(response.Body)				//used io.ReadAll	

	if err2 != nil {
		panic(err2)
	}
	content := string(databytes2)
	fmt.Println(content)
}
