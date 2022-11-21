package main

import (
	"fmt"
	"math/rand"
	"github.com/gammazero/deque"
)

func createRandomGraph(n, p int) map[int][]int {
	res := make(map[int][]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if rand.Intn(100) < p {
				list, ok := res[i]
				if !ok {
					list = make([]int,0)
				}
				list = append(list, j)
				res[i] = list

				list, ok = res[j]
				if !ok {
					list = make([]int,0)
				}
				list = append(list, i)
				res[j] = list
			}
		}
	}
	return res
}

type Elements struct {
	first int
	second int
}

func bfs(start, end int, graph map[int][]int) int {
	var queue deque.Deque[Elements]
	queue.PushBack(Elements{0, start})
	visited := make(map[int]bool)
	for queue.Len() != 0 {
		elem := queue.PopFront()
		steps := elem.first
		node := elem.second
		if node == end {
			return steps
		}
		for _, neigh := range graph[node] {
			_, seen := visited[neigh]
			if !seen {
				queue.PushBack(Elements{steps+1, neigh})
				visited[neigh] = true
			}
		}
	}
	return -1
}

func main() {
	n := 10000
	p := 10
	start := 0
	end := 999
	graph := createRandomGraph(n, p)
	steps := bfs(start, end, graph)
	fmt.Println("shortest distance between node", start, "and node", end, "is", steps)
}
