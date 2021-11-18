package graph

import "testing"

func TestNewGraph(t *testing.T) {
	expected := true
	v1 := &Vertex{value: 1}
	v2 := &Vertex{value: 2}
	v3 := &Vertex{value: 3}
	v4 := &Vertex{value: 4}
	e1 := &Edge{start: v1, end: v2, weight: 1}
	e2 := &Edge{start: v1, end: v3, weight: 2}
	e3 := &Edge{start: v2, end: v4, weight: 3}
	g1 := &Graph{vertecies: []*Vertex{v1, v2, v3, v4}, edges: []*Edge{e1, e2, e3}}
	graph := NewGraph()
	vertArr := []int{1, 2, 3, 4}
	vertecies := graph.AddVertecies(vertArr)
	graph.AddEdge(vertecies[0], vertecies[1], 1)
	graph.AddEdge(vertecies[0], vertecies[2], 2)
	graph.AddEdge(vertecies[1], vertecies[3], 3)
	result := graph.IsSame(g1)
	if expected != result {
		t.Errorf("FAILURE: New Graph not created correctly")
	} else {
		t.Logf("SUCCESS: New graph created correctly, got %s", graph.String())
	}
}

func TestShortestPath(t *testing.T) {
	expected := 4
	graph := NewGraph()
	vertArr := []int{0, 1, 2, 3}
	vertecies := graph.AddVertecies(vertArr)
	graph.AddEdge(vertecies[0], vertecies[1], 1)
	graph.AddEdge(vertecies[0], vertecies[2], 2)
	graph.AddEdge(vertecies[1], vertecies[3], 3)
	result, path := graph.FindShortestPath(vertecies[0], vertecies[3])
	if expected != result {
		t.Errorf("FAILURE: shortest path is %d but got %d", expected, result)
	} else {
		t.Logf("SUCCESS: shortest path found, got %s along the path %v", graph.String(), path)
	}
}

func TestShortestPath2(t *testing.T) {
	expected := 15
	graph := NewGraph()
	vertArr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	vertecies := graph.AddVertecies(vertArr)
	graph.AddEdge(vertecies[0], vertecies[1], 1)
	graph.AddEdge(vertecies[0], vertecies[2], 2)
	graph.AddEdge(vertecies[1], vertecies[3], 3)
	graph.AddEdge(vertecies[3], vertecies[4], 3)
	graph.AddEdge(vertecies[0], vertecies[5], 10)
	graph.AddEdge(vertecies[4], vertecies[5], 2)
	graph.AddEdge(vertecies[5], vertecies[6], 1)
	graph.AddEdge(vertecies[5], vertecies[8], 5)
	graph.AddEdge(vertecies[8], vertecies[9], 1)
	graph.AddEdge(vertecies[2], vertecies[7], 4)
	result, path := graph.FindShortestPath(vertecies[0], vertecies[9])
	if expected != result {
		t.Errorf("FAILURE: shortest path is %d but got %d", expected, result)
	} else {
		t.Logf("SUCCESS: shortest path found, got %s along the path %v", graph.String(), path)
	}
}
