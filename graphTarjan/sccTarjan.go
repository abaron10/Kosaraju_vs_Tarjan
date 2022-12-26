package graphTarjan

import (
	"fmt"
	"github.com/abaron10/Gothon/gothonSlice"
)

type GraphT struct {
	visitedNodes []int
	OnStack      map[int]bool
	stack        []int
	ids          map[int]int
	id           int
	low          map[int]int
	graphNodes   map[int][]*edge
}

type edge struct {
	from int
	to   int
}

func NewEdge(from int, to int) *edge {
	return &edge{from: from, to: to}
}

func NewGraphT() *GraphT {
	return &GraphT{visitedNodes: []int{}, ids: map[int]int{}, low: map[int]int{}, graphNodes: map[int][]*edge{}, OnStack: map[int]bool{}, stack: []int{}, id: -1}
}

func (g *GraphT) AddEdge(from int, to int) {
	edges, _ := g.graphNodes[from]
	edges = append(edges, NewEdge(from, to))

	g.graphNodes[from] = edges
	g.visitedNodes = append(g.visitedNodes, from)
}

func (g *GraphT) EvaluateTarjan() {
	visited := map[int]struct{}{}
	stack := []int{}

	for _, node := range g.visitedNodes {
		if _, alreadyVisitedNode := visited[node]; !alreadyVisitedNode {
			g.dfsImplementation(node, g.graphNodes, visited, &stack)
		}
	}

	//g.computeAnswer(g.low)
}

func (g *GraphT) computeAnswer(sccKeys map[int]int) {
	response := map[int][]int{}
	for node, sscComponent := range sccKeys {
		response[sscComponent] = append(response[sscComponent], node)
	}

	fmt.Println("The SCC (Strong connected components) calculated with Tarjans's Algorithm are:")
	for _, value := range response {
		fmt.Printf("- %v \n", value)
	}
}

func (g *GraphT) dfsImplementation(from int, graph map[int][]*edge, visited map[int]struct{}, stack *[]int) {
	visited[from] = struct{}{}
	g.id++
	g.ids[from] = g.id
	g.low[from] = g.id
	g.OnStack[from] = true
	g.stack = append(g.stack, from)

	edges, ok := graph[from]
	if ok {
		for _, edge := range edges {
			if _, visitedNode := visited[edge.to]; !visitedNode {
				g.dfsImplementation(edge.to, graph, visited, stack)
			}

			if _, isOnStack := g.OnStack[edge.to]; isOnStack {
				g.low[from] = Min(g.low[from], g.low[edge.to])
			}
		}
	}

	if g.ids[from] == g.low[from] {

		for true {
			node := gothonSlice.Pop(&g.stack, -1)
			delete(g.OnStack, node)
			g.low[node] = g.ids[from]

			if node == from {
				break
			}
		}
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
