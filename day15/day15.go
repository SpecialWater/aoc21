package day15

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/RyanCarrier/dijkstra"
)

var graph = dijkstra.NewGraph()
var graph2 = dijkstra.NewGraph()

func Day15() {
	data := readFile("day15/input.txt")
	BuildGraph(data, graph, 100)

	best, err := graph.Shortest(0, 9999)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Shortest distance ", best.Distance, " following path ", best.Path)

	data2 := expand(data)
	BuildGraph(data2, graph2, 500)
	best, err = graph2.Shortest(0, 249999)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Shortest distance ", best.Distance, " following path ", best.Path)
	//fmt.Printf("Answer to Day12, Q1: %v\n", ans1)
	//fmt.Printf("Answer to Day12, Q2: %v\n", ans2)
}

func expand(data map[string]int) map[string]int {
	data2 := make(map[string]int)
	counter := 0
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			for k, v := range data {
				coords := strings.Split(k, ",")
				coord_row, _ := strconv.Atoi(coords[0])
				coord_col, _ := strconv.Atoi(coords[1])
				coord := strconv.Itoa(coord_row+row*100) + "," + strconv.Itoa(coord_col+col*100)

				if newVal := v + row + col; newVal > 9 {
					newVal -= 9
					data2[coord] = newVal
				} else {
					data2[coord] = newVal
				}
				graph2.AddVertex(counter)
				counter += 1

			}
		}
	}

	return data2
}

func BuildGraph(coordinates map[string]int, g *dijkstra.Graph, size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			coord_right := strconv.Itoa(i) + "," + strconv.Itoa(j+1)
			coord_left := strconv.Itoa(i) + "," + strconv.Itoa(j-1)
			coord_up := strconv.Itoa(i-1) + "," + strconv.Itoa(j)
			coord_down := strconv.Itoa(i+1) + "," + strconv.Itoa(j)

			if weight, ok := coordinates[coord_right]; ok {
				g.AddArc(i*size+j, i*size+j+1, int64(weight))
			}

			if weight, ok := coordinates[coord_left]; ok {
				g.AddArc(i*size+j, i*size+j-1, int64(weight))
			}

			if weight, ok := coordinates[coord_up]; ok {
				g.AddArc(i*size+j, (i-1)*size+j, int64(weight))
			}

			if weight, ok := coordinates[coord_down]; ok {
				g.AddArc(i*size+j, (i+1)*size+j, int64(weight))
			}
		}
	}

}

func readFile(path string) map[string]int {
	file, err := os.Open(path)
	data := make(map[string]int)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	row := 0
	for scanner.Scan() {
		l := scanner.Text()
		for i, v := range l {
			val, _ := strconv.Atoi(string(v))
			coord := strconv.Itoa(row) + "," + strconv.Itoa(i)
			data[coord] = val

			graph.AddVertex(row*100 + i)
		}
		row += 1
	}
	return data

}
