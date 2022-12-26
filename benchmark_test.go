package main

import (
	"SCC_analysis/graphKosaraju"
	"SCC_analysis/graphTarjan"
	"testing"
)

func BenchmarkString(b *testing.B) {
	b.Run("tarjan", func(b *testing.B) {
		for i := 0; i < 100000; i++ {
			d := graphTarjan.NewGraphT()
			d.AddEdge(0, 1)
			d.AddEdge(0, 5)
			d.AddEdge(3, 5)
			d.AddEdge(3, 2)
			d.AddEdge(2, 0)
			d.AddEdge(2, 3)
			d.AddEdge(3, 3)
			d.AddEdge(5, 4)
			d.AddEdge(4, 3)
			d.AddEdge(4, 2)
			d.AddEdge(6, 4)
			d.AddEdge(6, 0)
			d.AddEdge(6, 8)
			d.AddEdge(8, 6)
			d.AddEdge(7, 6)
			d.AddEdge(6, 9)
			d.AddEdge(7, 9)
			d.AddEdge(11, 4)
			d.AddEdge(9, 11)
			d.AddEdge(9, 10)
			d.AddEdge(10, 12)
			d.AddEdge(12, 9)
			d.AddEdge(11, 12)

			d.EvaluateTarjan()

		}
	})
	b.Run("Kosaraju", func(b *testing.B) {
		for i := 0; i < 100000; i++ {
			d := graphKosaraju.NewGraphK()
			d.AddEdge(0, 1)
			d.AddEdge(0, 5)
			d.AddEdge(3, 5)
			d.AddEdge(3, 2)
			d.AddEdge(2, 0)
			d.AddEdge(2, 3)
			d.AddEdge(3, 3)
			d.AddEdge(5, 4)
			d.AddEdge(4, 3)
			d.AddEdge(4, 2)
			d.AddEdge(6, 4)
			d.AddEdge(6, 0)
			d.AddEdge(6, 8)
			d.AddEdge(8, 6)
			d.AddEdge(7, 6)
			d.AddEdge(6, 9)
			d.AddEdge(7, 9)
			d.AddEdge(11, 4)
			d.AddEdge(9, 11)
			d.AddEdge(9, 10)
			d.AddEdge(10, 12)
			d.AddEdge(12, 9)
			d.AddEdge(11, 12)

			d.EvaluateKosaraju()
		}
	})
}
