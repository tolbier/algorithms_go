package graphs

import (
	"fmt"
)

type Graph struct {
	isDirected  bool
	edges       Edges
	verticesMap VerticesMap
}

func NewGraph(isDirected bool) *Graph {
	return &Graph{
		isDirected:  isDirected,
		edges:       Edges{},
		verticesMap: VerticesMap{},
	}
}

func (g *Graph) AddEdge(id1, id2 int) {
	g.AddEdgeWithWeight(id1, id2, 0)

}
func (g *Graph) AddEdgeWithWeight(id1, id2 int, weight float64) {
	v1, v2 := g.verticesMap.getVertices(id1, id2)
	edge := NewEdge(v1, v2, g.isDirected, SetWeight(weight))
	g.edges = append(g.edges, edge)
	v1.AddAdjacentVertex(edge, v2)
	if g.IsUndirected() {
		v2.AddAdjacentVertex(edge, v1)
	}
}

func (g *Graph) IsDirected() bool {
	return g.isDirected
}

func (g *Graph) IsUndirected() bool {
	return !g.isDirected
}

func (g *Graph) getAllEdgesSorted() Edges {
	return g.edges.sorted()
}
func (g *Graph) getAllVertices() (vertices Vertices) {
	return g.verticesMap.getSortedVertices()
}

func (g *Graph) String() string {
	return fmt.Sprintf("{%v,%s}", g.getAllVertices(), g.getAllEdgesSorted())
}
