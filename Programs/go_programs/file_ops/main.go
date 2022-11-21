package main

import (
	"strconv"
	"math/rand"
	"io/ioutil"
)

func main() {
	// read random text file
	content, err := ioutil.ReadFile("random_text.txt")
	if err != nil {
		return
	}
	// write mutated files
	numMutations := 1000
	for i := 0; i < numMutations; i++ {
		// generate random number (representing character index)
		num := rand.Intn(len(content))
		// replace the random character by a
		newData := make([]byte, len(content))
		copy(newData, content)
		newData[num] = ([]byte("a"))[0]
		// write new data to file
		fileName := "random_text_mut_" + strconv.Itoa(i) + ".text"
		ioutil.WriteFile(fileName, newData, 0644)
	}
}
