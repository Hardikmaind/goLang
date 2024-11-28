//!Example 2: Using default to Avoid Blocking

package main

import "fmt"

func main() {
	channel := make(chan int)

	select {
	case msg := <-channel:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No messages available, continuing...")
	}
}
