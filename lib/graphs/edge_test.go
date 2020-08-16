package graphs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/tolbier/algorithms_go/lib/graphs"
)

func Test_NewEdge_Directed(t *testing.T) {
	e := NewEdge(NewVertex(1), NewVertex(5), true)
	assert.Equal(t, 1, e.GetVertex1().ID)
	assert.Equal(t, 5, e.GetVertex2().ID)
	assert.Equal(t, true, e.IsDirected())
	assert.Equal(t, false, e.IsUndirected())
}
func Test_NewEdge_Undirected(t *testing.T) {
	e := NewEdge(NewVertex(1), NewVertex(5), false)
	assert.Equal(t, 1, e.GetVertex1().ID)
	assert.Equal(t, 5, e.GetVertex2().ID)
	assert.Equal(t, false, e.IsDirected())
	assert.Equal(t, true, e.IsUndirected())

}
func Test_NewEdge_WithWeight_Directed(t *testing.T) {
	e := NewEdge(NewVertex(1), NewVertex(5), true, SetWeight(3))
	assert.Equal(t, 1, e.GetVertex1().ID)
	assert.Equal(t, 5, e.GetVertex2().ID)
	assert.Equal(t, true, e.IsDirected())
	assert.Equal(t, false, e.IsUndirected())
	assert.Equal(t, float64(3), e.Weight())
}
