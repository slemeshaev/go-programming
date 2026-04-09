package main

import "fmt"

type Graph struct {
	nodes map[string]map[string]bool
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]map[string]bool),
	}
}

func (g *Graph) AddEdge(from, to string) {
	if g.nodes[from] == nil {
		g.nodes[from] = make(map[string]bool)
	}

	g.nodes[from][to] = true
}

func (g *Graph) HasEdge(from, to string) bool {
	return g.nodes[from] != nil && g.nodes[from][to]
}

func main() {
	g := NewGraph()
	g.AddEdge("A", "B")

	fmt.Println(g.HasEdge("A", "B"))
	fmt.Println(g.HasEdge("B", "A"))
}
