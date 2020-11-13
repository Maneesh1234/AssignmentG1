package main

import (
	"fmt"
	"time"
)

// The pinger prints a ping and waits for a pong
func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		v1 := <-pinger
		fmt.Println("ping", v1)
		time.Sleep(time.Second)
		ponger <- 0
	}
}

// The ponger prints a pong and waits for a ping
func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		v1 := <-ponger
		fmt.Println("pong", v1)
		time.Sleep(time.Second)
		pinger <- 1
	}
}

func main() {
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	// The main goroutine starts the ping/pong by sending into the ping channel
	ping <- 1

	for {
		// Block the main thread until an interrupt
		time.Sleep(time.Second)
	}
}
