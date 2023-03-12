package main

type Graph struct {
	nodes map[int]*Node
}

type Node struct {
	num  int
	next []*Node
}

func (g *Graph) AddNode(num int) {
	if _, ok := g.nodes[num]; !ok {
		g.nodes[num] = &Node{num: num}
	}
}

func (g *Graph) AddEdge(from, to int) {
	g.AddNode(from)
	g.AddNode(to)
	g.nodes[from].next = append(g.nodes[from].next, g.nodes[to])
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[int]*Node)}
}
