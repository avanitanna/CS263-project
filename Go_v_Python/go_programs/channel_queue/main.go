package main

import (
	"fmt"
	"time"
	"math/rand"
)

func producer(res chan float64, done chan bool) {
	fmt.Println("running producer")
	for i := 0; i < 1000; i++ {
		// generate random value between 0 and 1
		v := rand.Float64()
		// block for that fraction of a second
		time.Sleep(time.Duration(v  * 1000 * 1000 * 1000) * time.Nanosecond)
		// add value to the queue
		res <- v
	}
	res <- -1
	fmt.Println("completed running producer")
	done <- true
}

func consumer(inp chan float64, done chan bool) {
	fmt.Println("running consumer")
	for true {
		// get value from the queue
		v := <-inp
		// if value is -1, bream from the infinite loop and this trhead will be terminated
		if v == -1 {
			break
		}
		// print the received value from the shared queue
		fmt.Println("Received valued", v)
	}
	fmt.Println("completed running consumer")
	done <- true
}

func main() {
	done := make(chan bool, 2)
	queue := make(chan float64)
	go consumer(queue, done)
	go producer(queue, done)
	<-done
	<-done
}
