package main

import (
	"fmt"
)

func worker(id int, job chan int, done chan int) {
	iterations := 0
	for true {
		cur := <-job
		if cur < 0 {
			break
		}
		fmt.Printf("[Thread %v] job: %v\n", id, cur)
		iterations++
	}
	done<-iterations
}

func main() {
	jobs := []int{1,2,3,4,5,6,7,8,9,9,8,7,6,5,4,3,2,1}
	num_threads := 4
	jobqueue := make(chan int, len(jobs))
	done := make(chan int, 1)

	for i := 0; i < num_threads; i++ {
		go worker(i, jobqueue, done)
	}

	for _, job := range jobs {
		jobqueue<-job
	}

	for i := 0; i < num_threads; i++ {
		jobqueue<- -1
	}

	for i := 0; i < num_threads; i++ {
		its := <-done
		fmt.Printf("One thread finished after %v iterations\n", its)
	}

}
