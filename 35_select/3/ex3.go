//! Example 3: Handling Multiple Sends and Receives

package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- 42
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- 24
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println("Received from channel1:", msg1)
		case msg2 := <-channel2:
			fmt.Println("Received from channel2:", msg2)
		}
	}
}
