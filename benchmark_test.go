package main

import (
	"SCC_analysis/graph"
	"SCC_analysis/graphKosaraju"
	"SCC_analysis/graphTarjan"
	"testing"
)

func BenchmarkSCCs(b *testing.B) {

	edges := map[int][]int{0: {1, 5}, 3: {5, 2, 3}, 2: {0, 3},
		5: {4}, 4: {3, 2}, 6: {4, 0, 8, 9}, 8: {6, 7},
		7: {6, 9}, 11: {4, 12}, 9: {11, 10}, 10: {12},
		12: {9}}

	b.Run("SCC with Tarjan's Algorithm", func(b *testing.B) {
		graphWithTarjan := graphTarjan.NewGraphT()
		graph.PopulateGraph(graphWithTarjan, edges)

		for i := 0; i < b.N; i++ {
			graphWithTarjan.EvaluateSCC()

		}
	})
	b.Run("SCC with Kosaraju's Algorithm", func(b *testing.B) {
		graphWithKosaraju := graphKosaraju.NewGraphK()
		graph.PopulateGraph(graphWithKosaraju, edges)

		for i := 0; i < b.N; i++ {
			graphWithKosaraju.EvaluateSCC()
		}
	})
}
