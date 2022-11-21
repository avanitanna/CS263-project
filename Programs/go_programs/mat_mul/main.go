package main

import (
	"math/rand"
)

func generateMat(matDim int) ([][]int, [][]int) {
	var resA, resB [][]int
	for i := 0; i < matDim; i++ {
		resA = append(resA, make([]int, matDim))
		resB = append(resA, make([]int, matDim))
		for j := 0; j < matDim; j++ {
			resA[i][j] = rand.Int()
			resB[i][j] = rand.Int()
		}
	}
	return resA, resB
}

func main() {
	numThreads := 5
	matDim := 100
	var matA, matB, res [][]int
	for i := 0; i < matDim; i++ {
		res = append(res, make([]int, matDim))
	}

	matA, matB = generateMat(matDim)

	done := make(chan bool)

	for i := 0; i < numThreads; i++ {
		go func(start, end int) {
			for j := start; j < end; j++ {
				for k := 0; k < matDim; k++ {
					for l := 0; l < matDim; l++ {
						res[j][k] += matA[j][l] * matB[l][k]
					}
				}
			}
			done <- true
		}(int(float64(matDim)/float64(numThreads) * float64(i)), int(float64(matDim)/float64(numThreads) * float64(i+1)))
	}
	for i := 0; i < numThreads; i++ {
		<-done
	}
}
