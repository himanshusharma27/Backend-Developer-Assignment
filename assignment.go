package main

import (
	"fmt"
	"math"
)

// Graph represents the transportation network
type Graph map[string]map[string]float64

// Dijkstra's algorithm to find the shortest path
func dijkstra(graph Graph, start, end string) (map[string]float64, map[string]string) {
	distances := make(map[string]float64)
	previous := make(map[string]string)
	unvisited := make(map[string]bool)

	// Initialize distances and unvisited nodes
	for node := range graph {
		distances[node] = math.Inf(1)
		previous[node] = ""
		unvisited[node] = true
	}
	distances[start] = 0

	for len(unvisited) > 0 {
		// Find the node with the smallest distance
		var current string
		minDist := math.Inf(1)
		for node := range unvisited {
			if distances[node] < minDist {
				minDist = distances[node]
				current = node
			}
		}

		// Remove current node from unvisited set
		delete(unvisited, current)

		// Stop if we reach the end node
		if current == end {
			break
		}

		// Update distances for neighboring nodes
		for neighbor, weight := range graph[current] {
			newDist := distances[current] + weight
			if newDist < distances[neighbor] {
				distances[neighbor] = newDist
				previous[neighbor] = current
			}
		}
	}

	return distances, previous
}

// Function to retrieve the shortest path
func getShortestPath(previous map[string]string, start, end string) []string {
	var path []string
	current := end

	// Reconstruct the path in reverse order
	for current != "" {
		path = append([]string{current}, path...)
		current = previous[current]
	}

	// Ensure the start node is included in the path
	if len(path) > 0 && path[0] != start {
		path = append([]string{start}, path...)
	}

	return path
}

func main() {
	// Define the transportation network graph
	graph := Graph{
		"Space Station": {
			"Hub A": 10,
			"Hub B": 5,
		},
		"Hub A": {
			"Hub B": 2,
			"Hub C": 1,
		},
		"Hub B": {
			"Hub C": 4,
			"Hub D": 8,
		},
		"Hub C": {
			"Hub D": 2,
			"Planet Surface": 6,
		},
		"Hub D": {
			"Planet Surface": 3,
		},
		"Planet Surface": {},
	}

	start := "Space Station"
	end := "Planet Surface"

	// Find the shortest path
	distances, previous := dijkstra(graph, start, end)
	shortestPath := getShortestPath(previous, start, end)

	// Display the shortest path and distances
	fmt.Printf("Shortest Path from %s to %s: %v\n", start, end, shortestPath)
	fmt.Println("Distances:", distances)
}
