package graph

import "testing"

func TestMartixGraph(t *testing.T) {
	edges := make([]Edge, 10)
	for i := 0; i < 10; i++ {
		e := Edge{}
		e.V1 = Vertex(i)
		e.V2 = Vertex(i)
		e.Weight = i
		edges = append(edges, e)
	}
	graph := BuildGraph(10, edges)
	graph.TraversalGraph()
}
