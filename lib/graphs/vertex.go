package graphs

import (
	"fmt"
	"sort"
)

type VerticesMap map[int]*Vertex
type Vertices []*Vertex
type Vertex struct {
	ID               int
	Value            interface{}
	inEdges, edges   Edges
	adjacentVertices Vertices
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
		adjacentVertices: Vertices{},
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
func (v *Vertex) GetAdjacentVertices() Vertices {
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
func (vm VerticesMap) getSortedVertices() (vertices Vertices) {
	vertices = make(Vertices, 0, len(vm))

	for _, vertex := range vm {
		vertices = append(vertices, vertex)
	}
	sort.Sort(vertices)
	return
}
func (vm VerticesMap) getVertices(id1, id2 int) (v1, v2 *Vertex) {
	return vm.getVertex(id1), vm.getVertex(id2)
}
func (vm VerticesMap) getVertex(id int) (v *Vertex) {
	var ok bool
	if v, ok = vm[id]; !ok {
		v = NewVertex(id)
		vm[v.ID] = v
	}
	return v
}
func (vv Vertices) Len() int { return len(vv) }
func (vv Vertices) Less(i, j int) bool {
	return vv[i].ID < vv[j].ID
}
func (vv Vertices) Swap(i, j int) { vv[i], vv[j] = vv[j], vv[i] }
