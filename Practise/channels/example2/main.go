package main
import "fmt"

func partialArraySum(start, step int, array []int, res chan int) {
	sum := 0
	for i := start; i < len(array); i += step {
		sum += array[i]
	}
	res<- sum
}

func main() {
	array := make([]int, 100)
	for i := range array {
		array[i] = i + 1
	}
	numThreads := 4
	res := make(chan int, 4)
	for i := 0; i < numThreads; i++ {
		go partialArraySum(i, numThreads, array, res)
	}
	sum := 0
	for i := 0; i < numThreads; i++ {
		sum += <-res
	}
	fmt.Println("Result:", sum)
}
