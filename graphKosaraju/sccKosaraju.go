package graphKosaraju

import (
	"fmt"
	"github.com/abaron10/Gothon/gothonSlice"
)

type GraphK struct {
	visitedNodes  []int
	reversedGraph *GraphK
	graphNodes    map[int][]*edge
}

type edge struct {
	from int
	to   int
}

func NewEdge(from int, to int) *edge {
	return &edge{from: from, to: to}
}

func NewGraphK() *GraphK {
	reversedGraph := &GraphK{visitedNodes: []int{}, graphNodes: map[int][]*edge{}}
	return &GraphK{visitedNodes: []int{}, graphNodes: map[int][]*edge{}, reversedGraph: reversedGraph}
}

func (g *GraphK) AddEdge(from int, to int) {
	if g == nil {
		return
	}
	edges, _ := g.graphNodes[from]
	edges = append(edges, NewEdge(from, to))

	g.graphNodes[from] = edges
	g.visitedNodes = append(g.visitedNodes, from)
	//Computing transposed graph at the same time user adds an Edge.
	g.reversedGraph.AddEdge(to, from)
}

func (g *GraphK) EvaluateKosaraju() {
	baseOrder := g.calculateStackBaseOrder()
	g.findSCCComponents(baseOrder)
	//g.computeAnswer(r)
}

func (g *GraphK) calculateStackBaseOrder() []int {
	visited := map[int]struct{}{}
	stack := []int{}

	for _, node := range g.visitedNodes {
		if _, alreadyVisitedNode := visited[node]; !alreadyVisitedNode {
			g.dfsImplementation(node, visited, &stack)
		}
	}

	return stack
}

func (g *GraphK) findSCCComponents(orderNodes []int) map[int]int {
	visited := map[int]struct{}{}
	SCC := map[int]int{}
	counter := 0

	for len(orderNodes) > 0 {
		//using https://github.com/abaron10/Gothon library to emulate stack popping as Python do
		node := gothonSlice.Pop(&orderNodes, -1)
		if _, alreadyVisitedNode := visited[node]; !alreadyVisitedNode {
			g.dfsImplementationReversed(node, visited, SCC, counter)
			counter++
		}
	}

	return SCC
}

/*
---------------------------------------------------------------------------------------------------------
--------------------------------DFS IMPLEMENTATIONS (BASE AND TRANSPOSED GRAPH-----------------------------
-----------------------------------------------------------------------------------------------------------
*/
func (g *GraphK) dfsImplementation(from int, visited map[int]struct{}, stack *[]int) {
	visited[from] = struct{}{}
	edges, ok := g.graphNodes[from]
	if ok {
		for _, edge := range edges {
			if _, visitedNode := visited[edge.to]; !visitedNode {
				g.dfsImplementation(edge.to, visited, stack)
			}
		}
	}
	*stack = append(*stack, from)
}

func (g *GraphK) dfsImplementationReversed(from int, visited map[int]struct{}, SCC map[int]int, counter int) {
	visited[from] = struct{}{}
	SCC[from] = counter
	edges, ok := g.reversedGraph.graphNodes[from]
	if ok {
		for _, edge := range edges {
			if _, visitedNode := visited[edge.to]; !visitedNode {
				g.dfsImplementationReversed(edge.to, visited, SCC, counter)
			}
		}
	}
	SCC[from] = counter
}

/*---------------------------------------------------------------------------------------------------------
--------------------------------HELPER METHOD TO COMPUTE A READABLE SCC ANSWER-----------------------------
-----------------------------------------------------------------------------------------------------------
*/

func (g *GraphK) computeAnswer(sccKeys map[int]int) {
	response := map[int][]int{}
	for node, sscComponent := range sccKeys {
		response[sscComponent] = append(response[sscComponent], node)
	}

	fmt.Println("The SCC (Strong connected components) calculated with Kosaraju's Algorithm are:")
	for _, value := range response {
		fmt.Printf("- %v \n", value)
	}
}
