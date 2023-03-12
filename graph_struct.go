package main

// Defining Graph structure as a map of nodes.
type Graph struct {
	nodes map[int]*Node
}

// Each course will be represented as a node
type Node struct {
	num  int
	next []*Node
}

// Track the visited and completed status
type Track struct {
	visited   bool
	completed bool
}

// AddNode function adds a new node to the given graph with given data
func (g *Graph) AddNode(num int) {
	if _, ok := g.nodes[num]; !ok {
		g.nodes[num] = &Node{num: num}
	}
}

// AddEdge function adds a relationship between given node and target node
func (g *Graph) AddEdge(from, to int) {
	g.AddNode(from)
	g.AddNode(to)
	g.nodes[from].next = append(g.nodes[from].next, g.nodes[to])
}

// NewGraph function initializes a new graph
func NewGraph() *Graph {
	return &Graph{nodes: make(map[int]*Node)}
}

// createGraph function creates a graph for the given prereqs and courses
func createGraph(n int, prereqs [][]int) *Graph {
	graph := NewGraph()

	for _, edge := range prereqs {
		graph.AddEdge(edge[1], edge[0])
	}

	return graph
}
