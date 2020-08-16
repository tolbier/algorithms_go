package graphs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tolbier/algorithms_go/lib/graphs"
)

func Test_NewGraph_Directed(t *testing.T) {
	g := graphs.NewGraph(true)
	for i := 5; i >= 1; i-- {
		for j := 5; j >= 1; j-- {
			if i != j {
				g.AddEdgeWithWeight(i, j, float64(j))
			}
		}
	}
	assert.Equal(t, true, g.IsDirected())
	assert.Equal(t, "{[{1(1):[2 3 4 5]} {2(2):[1 3 4 5]} {3(3):[1 2 4 5]} {4(4):[1 2 3 5]} {5(5):[1 2 3 4]}],[{1:2:2.00:D} {1:3:3.00:D} {1:4:4.00:D} {1:5:5.00:D} {2:1:1.00:D} {2:3:3.00:D} {2:4:4.00:D} {2:5:5.00:D} {3:1:1.00:D} {3:2:2.00:D} {3:4:4.00:D} {3:5:5.00:D} {4:1:1.00:D} {4:2:2.00:D} {4:3:3.00:D} {4:5:5.00:D} {5:1:1.00:D} {5:2:2.00:D} {5:3:3.00:D} {5:4:4.00:D}]}", g.String())
}
func Test_NewGraph_Directed_NoWeight(t *testing.T) {
	g := graphs.NewGraph(true)
	for i := 5; i >= 1; i-- {
		for j := 5; j >= 1; j-- {
			if i != j {
				g.AddEdge(i, j)
			}
		}
	}
	assert.Equal(t, true, g.IsDirected())
	assert.Equal(t, "{[{1(1):[2 3 4 5]} {2(2):[1 3 4 5]} {3(3):[1 2 4 5]} {4(4):[1 2 3 5]} {5(5):[1 2 3 4]}],[{1:2:0.00:D} {1:3:0.00:D} {1:4:0.00:D} {1:5:0.00:D} {2:1:0.00:D} {2:3:0.00:D} {2:4:0.00:D} {2:5:0.00:D} {3:1:0.00:D} {3:2:0.00:D} {3:4:0.00:D} {3:5:0.00:D} {4:1:0.00:D} {4:2:0.00:D} {4:3:0.00:D} {4:5:0.00:D} {5:1:0.00:D} {5:2:0.00:D} {5:3:0.00:D} {5:4:0.00:D}]}", g.String())
}
func Test_NewGraph_Undirected(t *testing.T) {
	g := graphs.NewGraph(false)
	for i := 5; i > 1; i-- {
		for j := i - 1; j >= 1; j-- {
			g.AddEdgeWithWeight(i, j, float64(j))
		}
	}
	assert.True(t, g.IsUndirected())
	assert.Equal(t, "{[{1(1):[2 3 4 5]} {2(2):[1 3 4 5]} {3(3):[1 2 4 5]} {4(4):[1 2 3 5]} {5(5):[1 2 3 4]}],[{2:1:1.00:U} {3:1:1.00:U} {3:2:2.00:U} {4:1:1.00:U} {4:2:2.00:U} {4:3:3.00:U} {5:1:1.00:U} {5:2:2.00:U} {5:3:3.00:U} {5:4:4.00:U}]}", g.String())
}
func Test_NewGraph_Undirected_NoWeight(t *testing.T) {
	g := graphs.NewGraph(false)
	for i := 5; i > 1; i-- {
		for j := i - 1; j >= 1; j-- {
			g.AddEdge(i, j)
		}
	}
	assert.True(t, g.IsUndirected())
	assert.Equal(t, "{[{1(1):[2 3 4 5]} {2(2):[1 3 4 5]} {3(3):[1 2 4 5]} {4(4):[1 2 3 5]} {5(5):[1 2 3 4]}],[{2:1:0.00:U} {3:1:0.00:U} {3:2:0.00:U} {4:1:0.00:U} {4:2:0.00:U} {4:3:0.00:U} {5:1:0.00:U} {5:2:0.00:U} {5:3:0.00:U} {5:4:0.00:U}]}", g.String())
}
