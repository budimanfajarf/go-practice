// Depth-first search (DFS)
// DFS use LIFO (Last-In, First-Out)
// https://youtu.be/Urx87-NMm6c?si=E-UnHFVzVmtRAzLq
// https://github.com/msambol/dsa/blob/master/search/depth_first_search.py

package main

import "fmt"

func dfs(graph map[string][]string, node string) {
	visited := map[string]bool{}
	stack := []string{}

	visited[node] = true
	stack = append(stack, node)

	fmt.Print("initial stack:", stack, " | ")
	fmt.Println("initial visited:", visited)

	for len(stack) > 0 {
		// Pop from stack
		lastIndex := len(stack) - 1
		s := stack[lastIndex]
		stack = stack[:lastIndex]

		fmt.Print("popped:", s, " | ")
		fmt.Print("stack now:", stack, " | ")

		// Reverse iterate through the edge list so results match recursive version.
		neighbors := graph[s]
		fmt.Print("neighbors:", neighbors, " | ")
		fmt.Print("visited:")
		for i := len(graph[s]) - 1; i >= 0; i-- {
			n := neighbors[i]
			if !visited[n] {
				visited[n] = true
				stack = append(stack, n)
				fmt.Print(n, ",")
			}
		}
		fmt.Print(" | ")

		fmt.Print("stack after:", stack, "\n")
	}

	fmt.Print("final stack:", stack, " | ")
	fmt.Println("final visited:", visited)
}

func main() {
	graph := map[string][]string{
		"A": {"B", "G"},
		"B": {"C", "D", "E"},
		"C": {},
		"D": {},
		"E": {"F"},
		"F": {},
		"G": {"H"},
		"H": {"I"},
		"I": {},
	}
	dfs(graph, "A")
}
