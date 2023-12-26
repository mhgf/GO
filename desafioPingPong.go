package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(1 * time.Millisecond)
			c1 <- "Pong"
		}
	}()

	go func() {
		for {
			time.Sleep(2 * time.Millisecond)
			c2 <- "Ping"
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

}
