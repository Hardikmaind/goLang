package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)

	// Goroutine to send a message after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "Hello from channel1"
	}()

	select {
	/*
		?below syntax means:
		!If channel1 sends a value, execute this case.
		!You don't capture or use the received value in this specific syntax.*/
	case <-channel1: // Just acknowledges receipt, but doesn't capture the value
		fmt.Println("A message was received on channel1")
	}
}

/*Anatomy of the Syntax:
1.channel1:
This is the channel you’re working with.
A channel is a conduit in Go that allows goroutines to communicate by passing data.

2.<-channel1:
The <- operator is used to receive a value from the channel.
It waits until channel1 sends a value, then "reads" that value.

3.case <-channel1::
Inside a select block, this means: "If channel1 is ready to send data, proceed with this case."
This doesn’t assign the received value to a variable; it just acknowledges the receipt.*/

/*
!Illustration Without select
If you were to use channel1 directly (outside select), receiving a value looks like this:

msg := <-channel1
fmt.Println("Received:", msg)

Here, msg is assigned the value sent through channel1.


!How It Works in a select Statement
When you write:

select {
case <-channel1:
    fmt.Println("Received a value from channel1")
}

It means:
If channel1 sends a value, execute this case.
You don't capture or use the received value in this specific syntax.

!Capturing the Received Value in select
If you want to capture the value sent by the channel, use:

select {
case msg := <-channel1:
    fmt.Println("Received:", msg)
}

Here:
msg is assigned the value sent by channel1.

!Summary
?case <-channel1:: Waits for channel1 to send data, but doesn’t use or capture the value.
?case msg := <-channel1:: Waits for channel1 to send data and stores the received value in msg.
?select ensures only one of the ready channels gets processed in this way.

This syntax is especially useful when you care about the channel event but not the actual value being sent.
*/
