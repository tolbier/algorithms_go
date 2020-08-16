package graphs

import (
	"fmt"
	"sort"
)

type Graph struct {
	isDirected  bool
	edges       Edges
	verticesMap map[int]*Vertex
}

func NewGraph(isDirected bool) *Graph {
	return &Graph{
		isDirected:  isDirected,
		edges:       Edges{},
		verticesMap: map[int]*Vertex{},
	}
}

func (g *Graph) AddEdge(id1, id2 int) {
	g.AddEdgeWithWeight(id1, id2, 0)

}
func (g *Graph) AddEdgeWithWeight(id1, id2 int, weight float64) {
	v1 := g.getVertex(id1)
	v2 := g.getVertex(id2)
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

func (g *Graph) GetAllEdgesSorted() Edges {
	sortedEdges := make(Edges, len(g.edges))
	copy(sortedEdges, g.edges)

	sort.SliceStable(sortedEdges, func(i, j int) bool {
		ei := sortedEdges[i]
		ej := sortedEdges[j]
		return ei.vertex1.ID < ej.vertex1.ID ||
			(ei.vertex1.ID == ej.vertex1.ID && ei.vertex2.ID < ej.vertex2.ID)
	})
	return sortedEdges
}
func (g *Graph) GetAllVertices() (vertices []*Vertex) {
	vertices = make([]*Vertex, 0, len(g.verticesMap))

	for _, vertex := range g.verticesMap {
		vertices = append(vertices, vertex)
	}
	sort.SliceStable(vertices, func(i, j int) bool {
		return vertices[i].ID < vertices[j].ID
	})
	return
}

func (g *Graph) getVertex(id1 int) (v *Vertex) {
	var ok bool
	if v, ok = g.verticesMap[id1]; !ok {
		v = NewVertex(id1)
		g.verticesMap[v.ID] = v
	}
	return v
}
func (g *Graph) String() string {
	return fmt.Sprintf("{%v,%s}", g.GetAllVertices(), g.GetAllEdgesSorted())
}
