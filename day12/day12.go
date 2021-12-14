package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func NewGraph() Graph {
	return Graph{
		adjacency: make(map[string][]string),
	}
}

type Graph struct {
	adjacency map[string][]string
}

func (g *Graph) AddVertex(vertex string) bool {
	if _, ok := g.adjacency[vertex]; ok {
		fmt.Printf("vertex %v already exists\n", vertex)
		return false
	}
	g.adjacency[vertex] = []string{}
	return true
}

func (g *Graph) AddEdge(vertex, node string) bool {
	if _, ok := g.adjacency[vertex]; !ok {
		fmt.Printf("vertex %v does not exists\n", vertex)
		return false
	}
	if ok := contains(g.adjacency[vertex], node); ok {
		fmt.Printf("node %v already exists\n", node)
		return false
	}

	g.adjacency[vertex] = append(g.adjacency[vertex], node)
	return true
}

var g = NewGraph()
var Visits = 0
var smallCaves = make(map[string]bool)

func Day12() {
	readFile("day12/input.txt")
	ans1, ans2 := g.DFS("start")

	fmt.Printf("Answer to Day12, Q1: %v\n", ans1)
	fmt.Printf("Answer to Day12, Q2: %v\n", ans2)

}

func (g Graph) DFS(startingNode string) (int, int) {
	visited := g.createVisited()
	g.dfsRecursive(startingNode, visited, "", false)
	ogVisits := Visits
	fmt.Println("ogVisits: ", ogVisits)
	newVisits := 0

	fmt.Println(smallCaves)
	for k, v := range smallCaves {
		Visits = 0
		g.dfsRecursive(startingNode, visited, k, v)
		newVisits += Visits - ogVisits
	}

	newVisits += ogVisits

	return ogVisits, newVisits
}

func (g Graph) dfsRecursive(startingNode string, visited map[string]bool, smallCave string, smallCaveVisited bool) {
	newVis := g.createVisited()
	for k, v := range visited {
		newVis[k] = v
	}

	if IsLower(startingNode) {
		newVis[startingNode] = true
	}

	if (startingNode == smallCave) && (!smallCaveVisited) {
		//fmt.Printf("Logic 1: %v, Logic 2: %v \n", startingNode == smallCave, !smallCaveVisited)
		newVis[startingNode] = false
		smallCaveVisited = true
	}

	if startingNode == "end" {
		Visits += 1
	}

	for _, node := range g.adjacency[startingNode] {

		if !newVis[node] {
			g.dfsRecursive(node, newVis, smallCave, smallCaveVisited)
		}

	}

}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func readFile(path string) {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	for scanner.Scan() {
		vertices := strings.Split(scanner.Text(), "-")
		g.AddVertex(vertices[0])
		g.AddVertex(vertices[1])
		g.AddEdge(vertices[0], vertices[1])
		g.AddEdge(vertices[1], vertices[0])

		if IsLower(vertices[0]) && (vertices[0] != "start" && vertices[0] != "end") {
			smallCaves[vertices[0]] = false
		}
		if IsLower(vertices[1]) && (vertices[1] != "start" && vertices[1] != "end") {
			smallCaves[vertices[1]] = false
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func (g Graph) createVisited() map[string]bool {
	visited := make(map[string]bool, len(g.adjacency))
	for key := range g.adjacency {
		visited[key] = false
	}
	return visited
}
