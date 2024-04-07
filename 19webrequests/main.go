package main

import (
	"fmt"
	// "io/ioutil"			//this is deprecreated, instead use io.ReadAll for reading the data
	"net/http"
	"io"
)

const url = "https://official-joke-api.appspot.com/ransdom_joke"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)
	fmt.Println("this is the err which has occured=>>>>>>>>",err)
	if err != nil {
// 		If the lines are not printing, it could be due to the buffering behavior of the standard output. By default, the standard output (fmt.Println) is line-buffered, meaning that it flushes its buffer and prints immediately when it encounters a newline character (\n). However, panic does not wait for buffering and immediately terminates the program.

// To force the lines to print before panicking, you can manually flush the output buffer

// this below lines will not print the error message, because the program is terminated before the buffer is flushed
		fmt.Println(err)
		fmt.Print("================================================")
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
