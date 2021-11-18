package graph

import (
	"fmt"
	"strings"
)

type Vertex struct {
	value int
}

type Edge struct {
	start  *Vertex
	end    *Vertex
	weight int
}

type Graph struct {
	vertecies []*Vertex
	edges     []*Edge
}

func NewGraph() *Graph {
	return &Graph{vertecies: make([]*Vertex, 0), edges: make([]*Edge, 0)}
}

func (g *Graph) AddVertecies(values []int) []*Vertex {
	for i := 0; i < len(values); i++ {
		g.vertecies = append(g.vertecies, &Vertex{value: values[i]})
	}
	return g.vertecies
}

func (g *Graph) AddEdge(s *Vertex, e *Vertex, w int) *Edge {
	if s == nil || e == nil {
		return nil
	}
	newEdge := &Edge{start: s, end: e, weight: w}
	g.edges = append(g.edges, newEdge)
	return newEdge
}

func (g *Graph) String() string {
	var output strings.Builder

	fmt.Fprint(&output, "Vertecies: ")
	for i := 0; i < len(g.vertecies); i++ {
		if i == 0 {
			fmt.Fprintf(&output, "%d", g.vertecies[i].value)
		} else {
			fmt.Fprintf(&output, ", %d", g.vertecies[i].value)
		}
		if i == len(g.vertecies) {
			fmt.Fprintf(&output, "\n")
		}
	}

	fmt.Fprint(&output, "Edges: ")
	for i := 0; i < len(g.edges); i++ {
		if i == 0 {
			fmt.Fprintf(&output, "%d -> %d (%d)", g.edges[i].start.value, g.edges[i].end.value, g.edges[i].weight)
		} else {
			fmt.Fprintf(&output, ", %d -> %d (%d)", g.edges[i].start.value, g.edges[i].end.value, g.edges[i].weight)
		}
		if i == len(g.edges) {
			fmt.Fprint(&output, "\n")
		}
	}

	return output.String()
}

func (g Graph) IsSame(g2 *Graph) bool {
	if len(g.vertecies) != len(g2.vertecies) || len(g.edges) != len(g2.edges) {
		return false
	}
	for i := 0; i < len(g.vertecies); i++ {
		if g.vertecies[i].value != g2.vertecies[i].value {
			return false
		}
	}
	for i := 0; i < len(g.edges); i++ {
		if g.edges[i].start.value != g2.edges[i].start.value ||
			g.edges[i].end.value != g2.edges[i].end.value {
			return false
		}
	}
	return true
}

func (g *Graph) FindShortestPath(start *Vertex, target *Vertex) (int, []int) {
	distMap, pathMap := g.CreateShortestPathDictionary(start, target)
	return distMap[target], pathMap[target]
}

func (g *Graph) CreateShortestPathDictionary(start *Vertex, target *Vertex) (map[*Vertex]int, map[*Vertex][]int) {
	visited := make(map[*Vertex]bool)
	nonVisited := make([]*Vertex, 0)
	distMap := make(map[*Vertex]int)
	pathMap := make(map[*Vertex][]int)

	nonVisited = append(nonVisited, start)
	distMap[start] = 0
	for {
		curr := closestNeighbor(nonVisited, distMap)
		for i := 0; i < len(g.edges); i++ {
			start := g.edges[i].start
			end := g.edges[i].end
			weight := g.edges[i].weight
			var neighbor *Vertex
			if start == curr {
				neighbor = end
			} else if end == curr {
				neighbor = start
			}
			if neighbor != nil {
				if !visited[neighbor] && !contains(nonVisited, neighbor) {
					nonVisited = append(nonVisited, end)
				}
				_, exists := distMap[neighbor]
				if !exists || distMap[neighbor] > distMap[curr]+weight {
					distMap[neighbor] = distMap[curr] + weight
					pathMap[neighbor] = append(pathMap[curr], curr.value)
				}
			}
		}
		visited[curr] = true
		nonVisited = remove(nonVisited, curr)
		if len(nonVisited) == 0 {
			break
		}
	}
	return distMap, pathMap
}

func closestNeighbor(nv []*Vertex, dm map[*Vertex]int) *Vertex {
	var closestVal *Vertex
	for i, neighbor := range nv {
		val, exists := dm[neighbor]
		if exists && i == 0 || val < dm[closestVal] {
			closestVal = neighbor
		}
	}
	return closestVal
}

func remove(arr []*Vertex, i *Vertex) []*Vertex {
	result := make([]*Vertex, 0)
	for j := 0; j < len(arr); j++ {
		if arr[j] != i {
			result = append(result, arr[j])
		}
	}
	return result
}

func contains(arr []*Vertex, target *Vertex) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return true
		}
	}
	return false
}
