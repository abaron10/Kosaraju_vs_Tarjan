package graph

type Graph interface {
	AddEdge(from int, to int)
	EvaluateSCC()
}

func PopulateGraph(graphImpl Graph, graph map[int][]int) {
	for node, edges := range graph {
		for _, edge := range edges {
			graphImpl.AddEdge(node, edge)
		}
	}
}
