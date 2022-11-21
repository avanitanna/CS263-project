package main

import (
	"log"
	"time"
	"math/rand"
)

func loggingRoutine(p float64, threadNum int, l *log.Logger, done chan bool) {
	l.Printf("[%v] running thread\n", threadNum)
	for i := 0; i < 100; i++ {
		// generate random value between 0 and 1
		v := rand.Float64()
		// log values
		l.Printf("[%v] Received value %v\n", threadNum, v)
		// block for that fraction of a second
		time.Sleep(time.Duration(v * 1000 * 1000 * 1000) * time.Nanosecond)
		// creates warning and stops is below a certain value p
		if v < p {
			l.Printf("[%v] Warning - we have reached a value less than %v\n", threadNum, p)
			break
		}
	}
	l.Printf("[%v] completed running thread\n", threadNum)
	done<- true
}

func main() {
	logger := log.Default()
	done := make(chan bool)
	threadNum := 10
	for i := 0; i < threadNum; i++ {
		go loggingRoutine(0.2, i, logger, done)
	}

	for i := 0; i < threadNum; i++ {
		<-done
	}
}
