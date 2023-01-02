package graphTarjan

import (
	"SCC_analysis/graph"
	"fmt"
	"github.com/abaron10/Gothon/gothonSlice"
)

type graphT struct {
	addedNodes []int
	onStack    map[int]struct{}
	stack      []int
	ids        map[int]int
	id         int
	lowLink    map[int]int
	graphNodes map[int][]*edge
}

type edge struct {
	from int
	to   int
}

func newEdge(from int, to int) *edge {
	return &edge{from: from, to: to}
}

func NewGraphT() graph.Graph {
	return &graphT{addedNodes: []int{}, ids: map[int]int{}, lowLink: map[int]int{}, graphNodes: map[int][]*edge{}, onStack: map[int]struct{}{}, stack: []int{}, id: -1}
}

func (g *graphT) AddEdge(from int, to int) {
	edges, _ := g.graphNodes[from]
	edges = append(edges, newEdge(from, to))

	g.graphNodes[from] = edges
	g.addedNodes = append(g.addedNodes, from)
}

func (g *graphT) EvaluateSCC() {
	visited := map[int]struct{}{}

	for _, node := range g.addedNodes {
		if _, alreadyVisitedNode := visited[node]; !alreadyVisitedNode {
			g.dfsImplementation(node, g.graphNodes, visited)
		}
	}
	//g.computeAnswer(g.lowLink)
}

func (g *graphT) computeAnswer(sccKeys map[int]int) {
	response := map[int][]int{}
	for node, sscComponent := range sccKeys {
		response[sscComponent] = append(response[sscComponent], node)
	}

	fmt.Println("The SCC (Strong connected components) calculated with Tarjan's Algorithm are:")
	for _, value := range response {
		fmt.Printf("- %v \n", value)
	}
}

func (g *graphT) dfsImplementation(from int, graph map[int][]*edge, visited map[int]struct{}) {
	visited[from] = struct{}{}
	g.id++
	g.ids[from] = g.id
	g.lowLink[from] = g.id

	g.onStack[from] = struct{}{}
	g.stack = append(g.stack, from)

	edges, ok := graph[from]
	if ok {
		for _, edge := range edges {
			if _, visitedNode := visited[edge.to]; !visitedNode {
				g.dfsImplementation(edge.to, graph, visited)
			}

			if _, isOnStack := g.onStack[edge.to]; isOnStack {
				g.lowLink[from] = Min(g.lowLink[from], g.lowLink[edge.to])
			}
		}
	}

	if g.ids[from] == g.lowLink[from] {

		for true {
			node := gothonSlice.Pop(&g.stack, -1)
			delete(g.onStack, node)
			g.lowLink[node] = g.ids[from]

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
