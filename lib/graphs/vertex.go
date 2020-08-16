package graphs

import (
	"fmt"
	"sort"
)

type Vertex struct {
	ID               int
	Value            interface{}
	inEdges, edges   Edges
	adjacentVertices []*Vertex
}

func NewVertex(ID int) *Vertex {
	return NewVertexWithValue(ID, ID)
}
func NewVertexWithValue(ID int, Value interface{}) *Vertex {
	return &Vertex{
		ID:               ID,
		Value:            Value,
		inEdges:          Edges{},
		edges:            Edges{},
		adjacentVertices: []*Vertex{},
	}
}

func (v *Vertex) AddAdjacentVertex(edge *Edge, adjVertex *Vertex) {
	v.edges = append(v.edges, edge)
	edge.vertex2.addInEdge(edge)
	v.adjacentVertices = append(v.adjacentVertices, adjVertex)
}

func (v *Vertex) addInEdge(edge *Edge) {
	v.inEdges = append(v.inEdges, edge)
}
func (v *Vertex) GetAdjacentVertices() []*Vertex {
	return v.adjacentVertices
}
func (v *Vertex) GetAdjacentIDs() (res []int) {
	for _, v := range v.GetAdjacentVertices() {
		res = append(res, v.ID)
	}
	sort.Ints(res)
	return
}

func (v *Vertex) String() string {
	return fmt.Sprintf("{%d(%v):%v}", v.ID, v.Value, v.GetAdjacentIDs())
}
