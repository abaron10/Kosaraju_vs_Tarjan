package graph

import (
	"fmt"
	"github.com/abaron10/Gothon/gothonSlice"
)

type Graph struct {
	visitedNodes  []int
	reversedGraph *Graph
	graphNodes    map[int][]*edge
}

type edge struct {
	from int
	to   int
}

func NewEdge(from int, to int) *edge {
	return &edge{from: from, to: to}
}

func NewGraph() *Graph {
	reversedGraph := &Graph{visitedNodes: []int{}, graphNodes: map[int][]*edge{}}
	return &Graph{visitedNodes: []int{}, graphNodes: map[int][]*edge{}, reversedGraph: reversedGraph}
}

func (g *Graph) AddEdge(from int, to int) {
	if g == nil {
		return
	}
	edges, _ := g.graphNodes[from]
	edges = append(edges, NewEdge(from, to))

	g.graphNodes[from] = edges
	g.visitedNodes = append(g.visitedNodes, from)
	g.reversedGraph.AddEdge(to, from)
}

func (g *Graph) EvaluateKosaraju() {
	calculateOrder := g.Dfs()

	r := g.DfsReversed(calculateOrder)
	fmt.Println(r)
}

func (g *Graph) Dfs() []int {
	visited := map[int]struct{}{}
	stack := []int{}

	for _, node := range g.visitedNodes {
		if _, alreadyVisitedNode := visited[node]; !alreadyVisitedNode {
			g.dfsImplementation(node, g.graphNodes, visited, &stack)
		}
	}

	return stack
}

func (g *Graph) DfsReversed(orderNodes []int) map[int]int {
	visited := map[int]struct{}{}
	SCC := map[int]int{}
	counter := 0

	for len(orderNodes) > 0 {
		node := gothonSlice.Pop(&orderNodes, -1)
		if _, alreadyVisitedNode := visited[node]; !alreadyVisitedNode {
			g.dfsImplementationReversed(node, g.reversedGraph.graphNodes, visited, SCC, counter)
			counter++
		}
	}

	return SCC
}

func (g *Graph) dfsImplementationReversed(from int, graph map[int][]*edge, visited map[int]struct{}, SCC map[int]int, counter int) {
	visited[from] = struct{}{}
	SCC[from] = counter
	edges, ok := graph[from]
	if ok {
		for _, edge := range edges {
			if _, visitedNode := visited[edge.to]; !visitedNode {
				g.dfsImplementationReversed(edge.to, graph, visited, SCC, counter)
			}
		}
	}
	SCC[from] = counter
}

func (g *Graph) dfsImplementation(from int, graph map[int][]*edge, visited map[int]struct{}, stack *[]int) {
	visited[from] = struct{}{}
	fmt.Println(from)
	edges, ok := graph[from]
	if ok {
		for _, edge := range edges {
			if _, visitedNode := visited[edge.to]; !visitedNode {
				g.dfsImplementation(edge.to, graph, visited, stack)
			}
		}
	}
	*stack = append(*stack, from)
}
