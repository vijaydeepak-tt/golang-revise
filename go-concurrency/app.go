package main

import (
	"fmt"
	"time"
)

func greet(msg string, done chan bool) {
	fmt.Println(msg)
	done <- true
}

func slowGreet(msg string, done chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello, ", msg)
	done <- true // <- represents the direction of data flow

	// need to close the channel in the last, else will get the error - fatal error: all goroutines are asleep - deadlock!
	// close function can be added only we know which one will completes last
	close(done)
}

// use separate channel for each process.
func main() {
	// dones := make([]chan bool, 4)
	done := make(chan bool)

	// dones[0] = make(chan bool)
	go greet("Msg 1", done)

	// dones[1] = make(chan bool)
	go greet("Msg 2", done)

	// dones[2] = make(chan bool)
	go slowGreet("Msg 3", done)

	// dones[3] = make(chan bool)
	go greet("Msg 4", done)

	// we are added 4 times the same parameter bcoz we using the same channel for 4 times.
	// if have separate channels then there will be 4 different line of code
	// Prints data in different order, based on the execution completion
	// <-done
	// <-done
	// <-done
	// <-done

	// for _, done := range dones {
	// 	<-done
	// }

	for range done {
	}
}
