package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// ! CONTEXT USE CASE IN DOING THE HTTP REQUEST
func main() {

	//create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Microsecond) //if time is 5 microsend we get context deadline exceeded error
	//if time is 5 seconds then we get the response status as OK 200
	defer cancel()

	//create a new request with context
	url := "https://timeapi.io/api/time/current/zone?timeZone=Europe%2FAmsterdam"

	//print the context when the context will TIMEOUT
	// Print the context deadline
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("Context deadline:", deadline)
	} else {
		fmt.Println("No deadline set")
	}

	//! method 1, to create a req with context
	/*
		req,_:=http.NewRequest("GET",url,nil)
		? make the http client
		client:=&http.Client{}
		resp,err:=client.Do(req.WithContext(ctx))
	*/

	//! method 2 , to create a request with a context
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	//	make the http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)               // we can also return the error returned by the context instead of the error returned by the http request
		fmt.Println("Context Error:", ctx.Err()) // this will print the error returned by the context
		return
	}
	defer resp.Body.Close()

	//check the response status
	fmt.Println("Response Status:", resp.Status)

}

//basically above code yeh kehta hain ki agar 5 microsecond ke andar response nahi aata toh accha hain, nahi toh hum cancel kar denge(context timeout jayega)
