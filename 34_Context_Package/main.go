package main

import (
    "context"
    "fmt"
    "time"
)

func demo(ctx context.Context, ch chan int) {
    time.Sleep(10 * time.Second)
    fmt.Println("10 seconds completed")
    ch <- 2
}

func main() {
    ch := make(chan int)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    go demo(ctx, ch)

    fmt.Println("Hello World")

    select {    
        /*
        ?below syntax means:
		!If ch sends a value, execute this case.
		!We don't capture or use the received value in this specific syntax. Similarly with the  ctx,Done() case below, which means "If the context is done, execute this"
        */
    case <-ch:          
        fmt.Println("Received from channel")
    case <-ctx.Done():              //this is  the channel context provides..to check if the context is done or not
        fmt.Println("Context timeout")
    }
}