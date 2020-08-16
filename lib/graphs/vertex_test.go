package graphs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/tolbier/algorithms_go/lib/graphs"
)

func Test_NewVertex(t *testing.T) {
	v := NewVertex(1)
	assert.Equal(t, 1, v.ID)
	assert.Equal(t, 1, v.Value)
}
func Test_NewVertexWithValue_With_Int(t *testing.T) {
	v := NewVertexWithValue(1, 5)
	assert.Equal(t, 1, v.ID)
	assert.Equal(t, 5, v.Value)
}
func Test_Vertex_With_String(t *testing.T) {
	v := NewVertexWithValue(1, "Hello")
	assert.Equal(t, 1, v.ID)
	assert.Equal(t, "Hello", v.Value)
}
