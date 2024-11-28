package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func openConnection(done chan bool) {
	fmt.Println("Attempting to open connection")

	if rand.Intn(100) > 50 {
		fmt.Println("OOPS hanging connection")
		time.Sleep(1000 * time.Hour)
	} else {
		time.Sleep(2 * time.Second)
		fmt.Println("Connection established")

	}
	done <- true //if rand.Intn(100) > 50 then this will not be executed till the time.Sleep(1000 * time.Hour) is completed and done channel will not be consumed , and it will waiting for the done channel to be consumed. but alongside our context is also running and it will cancel the context after 5 seconds and will print the message that connection took too long to establish and was cancelled
}
func openConnectionWithTimeout() {
	//ctx := context.Background() // ! This creates the simple context with no deadline no timeout nothing. type of ctx is context.Context()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // ! This creates the context with timeout of 5 seconds. type of ctx is context.Context()

	defer cancel() // ! This is important to call cancel function to release the resources.

	//? Run openConnection in a goroutine
	done := make(chan bool)
	go openConnection(done) //  ctx.Done() returns a channel which is closed when the context is done.

	select {
	case <-done:
		fmt.Println("openConnectionWithTimeout() - openConnection(),connection finished successfully")
	case <-ctx.Done(): //ctx.Done() returns a channel which is closed when the context is done. and this is the channel context provides..to check if the context is done or not
		fmt.Println("openConnectionWithTimeout() - openConnection() connection took too long to establish and was cancelled")
	}
}
func main() {
	openConnectionWithTimeout()

}
