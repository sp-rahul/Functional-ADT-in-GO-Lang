package main

import (
	"fmt"
)

// Graph represent an adjacency list graph
type Graph struct {
	vertices []vertex
}

// Vertex represent a graph
type vertex struct {
	key      int
	adjacent []vertex
}

//Add Edge

func (graph Graph) addEdge(from, to int) Graph {
	fromVertex := graph.getVertex(from)
	toVertex := graph.getVertex(to)
	fromVertex.adjacent = append(graph.getVertex(from).adjacent, toVertex)
	graph = graph.replaceAdacencyList(from, fromVertex)
	return graph
}

// Replace the new adjacent nodes list with the previous -> making immutable
func (graph Graph) replaceAdacencyList(k int, vertx vertex) Graph {
	for i, v := range graph.vertices {
		if v.key == k {
			graph.vertices[i].adjacent = vertx.adjacent
		}
	}
	return graph
}

// func (g Graph) addEdge(from, to int) Graph {
// 	//Get vertex
// 	fromVertex, err1 := g.getVertex(from)
// 	toVertex, err2 := g.getVertex(to)

// 	// Check error
// 	if err1 == false || err2 == false {
// 		err := fmt.Errorf("Invalid edge () %v --> %v )", from, to)
// 		fmt.Println(err.Error())

// 	} else if contains(fromVertex.adjacent, to) {

// 		err := fmt.Errorf("Already  edge () %v --> %v )", from, to)
// 		fmt.Println(err.Error())

// 	} else {

// 		// Add edge
// 		fromVertex = g.vertices[from]

// 		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)

// 	}
// 	return g
// }

// Get vertex
func (g Graph) getVertex(k int) vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}

	return g.vertices[0]

}

//addvertex adds a Vertex to the Graph
func (g Graph) addVertex(k int) Graph {

	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())

	} else {
		g.vertices = append(g.vertices, vertex{key: k})

	}
	return g
}

// printing the graph
func (g Graph) print() {
	for _, v := range g.vertices {

		fmt.Printf("\n Vertex %v: ", v.key)

		for _, u := range v.adjacent {
			fmt.Printf("%v ", u.key)
		}
	}
}

// vertiecs contaion or not check
func contains(s []vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}
func (g Graph) valSet(m, n int) {
	g.addVertex(n)
	if m == n {
		return
	}

	g.valSet(m, n-1)

}

func main() {
	test := Graph{}
	for i := 0; i < 5; i++ {
		test = test.addVertex(i)
	}
	//n := 5
	//test = test.valSet(0, n)

	test = test.addVertex(0)
	test = test.addVertex(4)

	test = test.addEdge(1, 2)
	test = test.addEdge(2, 3)
	test = test.addEdge(0, 2)
	test = test.addEdge(2, 4)
	test = test.addEdge(3, 3)
	// test = test.addEdge(6, 2)
	// test = test.addEdge(2, 4)
	// test = test.addEdge(6, 2)

	test.print()

}
