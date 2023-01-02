package graphKosaraju

import (
	"SCC_analysis/graph"
	"fmt"
	"github.com/abaron10/Gothon/gothonSlice"
)

type graphK struct {
	addedNodes    []int
	reversedGraph *graphK
	graphNodes    map[int][]*edge
}

type edge struct {
	from int
	to   int
}

func newEdge(from int, to int) *edge {
	return &edge{from: from, to: to}
}

func NewGraphK() graph.Graph {
	reversedGraph := &graphK{addedNodes: []int{}, graphNodes: map[int][]*edge{}}
	return &graphK{addedNodes: []int{}, graphNodes: map[int][]*edge{}, reversedGraph: reversedGraph}
}

func (g *graphK) AddEdge(from int, to int) {
	edges, _ := g.graphNodes[from]
	edges = append(edges, newEdge(from, to))

	g.graphNodes[from] = edges
	g.addedNodes = append(g.addedNodes, from)

	//Computing transposed graph at the same time user adds an Edge.
	if g.reversedGraph != nil {
		g.reversedGraph.AddEdge(to, from)
	}
}

func (g *graphK) EvaluateSCC() {
	baseOrder := g.calculateStackBaseOrder()
	g.findSCCComponents(baseOrder)
	//g.computeAnswer(sccResult)
}

func (g *graphK) calculateStackBaseOrder() []int {
	visited := map[int]struct{}{}
	stack := []int{}

	for _, node := range g.addedNodes {
		if _, alreadyVisitedNode := visited[node]; !alreadyVisitedNode {
			g.dfsImplementation(node, visited, &stack)
		}
	}

	return stack
}

func (g *graphK) findSCCComponents(orderNodes []int) map[int]int {
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
func (g *graphK) dfsImplementation(from int, visited map[int]struct{}, stack *[]int) {
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

func (g *graphK) dfsImplementationReversed(from int, visited map[int]struct{}, SCC map[int]int, counter int) {
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

func (g *graphK) computeAnswer(sccKeys map[int]int) {
	response := map[int][]int{}
	for node, sscComponent := range sccKeys {
		response[sscComponent] = append(response[sscComponent], node)
	}

	fmt.Println("The SCC (Strong connected components) calculated with Kosaraju's Algorithm are:")
	for _, value := range response {
		fmt.Printf("- %v \n", value)
	}
}
