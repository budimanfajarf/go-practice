// Bread-first search (BFS)
// BFS use FIFO (First-In, First-Out)
// hhttps://youtu.be/HZ5YTanv5QE?si=gHsuNEZPjOS7Q2ie
// https://github.com/msambol/dsa/blob/master/search/breadth_first_search.py

package main

import "fmt"

func bfs(graph map[string][]string, node string) {
	queue := []string{}
	visited := map[string]bool{}

	queue = append(queue, node)
	visited[node] = true

	fmt.Print("initial queue:", queue, " | ")
	fmt.Println("initial visited", visited)

	for len(queue) > 0 {
		// pop left from queue
		popped := queue[0]
		queue = queue[0+1:]

		fmt.Print("popped:", popped, " | ")
		fmt.Print("queue now:", queue, " | ")

		neighbors := graph[popped]
		fmt.Print("neighbors:", neighbors, " | ")
		fmt.Print("visited:")
		for _, n := range neighbors {
			if !visited[n] {
				queue = append(queue, n)
				visited[n] = true
				fmt.Print(n, ",")
			}
		}
		fmt.Print(" | ")

		fmt.Println("queue after", queue)
	}

	fmt.Print("final queue:", queue, " | ")
	fmt.Println("final visited:", visited)
}

func main() {
	var graph = map[string][]string{
		"A": {"B", "C"},
		"B": {"D", "E", "F"},
		"C": {"G"},
		"D": {},
		"E": {},
		"F": {"H"},
		"G": {"I"},
		"H": {},
		"I": {},
	}

	bfs(graph, "A")
}
