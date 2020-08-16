package graphs

import (
	"fmt"
	"sort"
)

type Edges []*Edge
type Directed bool
type Edge struct {
	vertex1    *Vertex
	vertex2    *Vertex
	isDirected Directed
	weight     float64
}

type EdgeOption func(*Edge) *Edge

func NewEdge(vertex1, vertex2 *Vertex, isDirected bool, options ...EdgeOption) (edge *Edge) {
	edge = &Edge{
		vertex1:    vertex1,
		vertex2:    vertex2,
		isDirected: Directed(isDirected),
	}
	for _, option := range options {
		edge = option(edge)
	}
	return
}

func SetWeight(weight float64) EdgeOption {
	return func(edge *Edge) *Edge {
		edge.weight = weight
		return edge
	}
}

func (e Edge) GetVertex1() *Vertex {
	return e.vertex1
}
func (e Edge) GetVertex2() *Vertex {
	return e.vertex2
}
func (e Edge) IsDirected() bool {
	return bool(e.isDirected)
}

func (e Edge) IsUndirected() bool {
	return !bool(e.isDirected)
}

func (e Edge) Weight() float64 {
	return e.weight
}
func (e Edge) String() string {
	return fmt.Sprintf("{%d:%d:%.2f:%s}", e.vertex1.ID, e.vertex2.ID, e.weight, e.isDirected)
}
func (et Directed) String() string {
	if et {
		return "D"
	}
	return "U"
}
func (ee Edges) String() string {
	res := fmt.Sprintf("%v", []*Edge(ee))
	return res
}

func (ee Edges) sorted() (sortedEdges Edges) {
	sortedEdges = make(Edges, len(ee))
	copy(sortedEdges, ee)
	sort.Sort(sortedEdges)
	return
}
func (ee Edges) Len() int { return len(ee) }
func (ee Edges) Less(i, j int) bool {
	return ee[i].vertex1.ID < ee[j].vertex1.ID ||
		(ee[i].vertex1.ID == ee[j].vertex1.ID && ee[i].vertex2.ID < ee[j].vertex2.ID)
}
func (ee Edges) Swap(i, j int) { ee[i], ee[j] = ee[j], ee[i] }
