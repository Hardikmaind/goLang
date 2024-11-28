// ! Example 1: Waiting for Multiple Channels

package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- "Message from channel 2"
	}()

	//? IF MULTIPLE CHANNELS ARE READY AT THE SAME TIME..THEN GO SELECTS THE RANDOM ONE.

	select {
	case msg1 := <-channel1:
		fmt.Println(msg1)
	case msg2 := <-channel2:
		fmt.Println(msg2)
	}
}
