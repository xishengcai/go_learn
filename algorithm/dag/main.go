package main

import (
	"fmt"
	"github.com/goombaio/dag"
)

func main() {
	// Create the dag
	dag1 := dag.NewDAG()

	// Create the vertices. Value is nil to simplify.
	vertex1 := dag.NewVertex("a", nil)
	vertex2 := dag.NewVertex("b", nil)
	vertex3 := dag.NewVertex("c", nil)
	vertex4 := dag.NewVertex("d", nil)

	// Add the vertices to the dag.
	dag1.AddVertex(vertex1)
	dag1.AddVertex(vertex2)
	dag1.AddVertex(vertex3)
	dag1.AddVertex(vertex4)

	// Add the edges (Note that given vertices must exist before adding an
	// edge between them).
	dag1.AddEdge(vertex1, vertex2)
	dag1.AddEdge(vertex2, vertex3)
	dag1.AddEdge(vertex2, vertex4)
	dag1.AddEdge(vertex4, vertex3)

	fmt.Println(dag1)
}
