package main

import (
	"fmt"
	"errors"
)

func binSearch(data []int, value int) (int, error) {
	upper, lower := len(data)-1, 0

	for lower <= upper {
		m := (lower + upper) / 2
		if data[m] < value {
			lower = m + 1
		} else if data[m] > value {
			upper = m - 1
		} else {
			return m, nil
		}
	}

	return 0, errors.New("Value not found")
}

func main() {

	data := []int{1,2,3,4,5,6,7,8,9}

	val := 2
	res, err := binSearch(data, val)
	if err == nil {
		fmt.Printf("Index of value %v: %v\n", val, res)
	} else {
		fmt.Printf("Could not find value %v: %v\n", val, err)
	}
}
