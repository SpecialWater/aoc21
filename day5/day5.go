package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	p1 Point
	p2 Point
}

func (line Line) minY() int {
	return min(line.p1.y, line.p2.y)
}
func (line Line) maxY() int {
	return max(line.p1.y, line.p2.y)
}
func (line Line) equalY() bool {
	return line.p1.y == line.p2.y
}

func (line Line) minX() int {
	return min(line.p1.x, line.p2.x)
}
func (line Line) maxX() int {
	return max(line.p1.x, line.p2.x)
}
func (line Line) equalX() bool {
	return line.p1.x == line.p2.x
}

func Day5() {
	lines := readFile("day5/input.txt")
	hvLines := horizontalVertical(lines)
	_, p1 := part1(hvLines)

	p2 := part2(lines)

	fmt.Printf("Answer to Day5, Q1: %v\n", p1)
	fmt.Printf("Answer to Day5, Q2: %v\n", p2)
}

func horizontalVertical(lines []Line) []Line {
	hvLines := []Line{}

	for _, pair := range lines {
		if pair.equalX() || pair.equalY() {
			hvLines = append(hvLines, pair)
		}
	}

	return hvLines
}

func part1(hvLines []Line) ([1000][1000]int, int) {
	grid := [1000][1000]int{}
	for _, pair := range hvLines {
		if pair.equalX() {
			for i := pair.minY(); i <= pair.maxY(); i++ {
				grid[pair.p1.x][i] += 1
			}
		}
		if pair.equalY() {
			for i := pair.minX(); i <= pair.maxX(); i++ {
				grid[i][pair.p1.y] += 1
			}
		}
	}

	return grid, overlap(grid)
}

func part2(hvDiagLines []Line) int {
	grid, _ := part1(hvDiagLines)
	for _, pair := range hvDiagLines {
		if !(pair.equalX() || pair.equalY()) {
			for i := 0; i <= (pair.maxX() - pair.minX()); i++ {
				p2BigY, p2BigX := (pair.p2.y-pair.p1.y) > 0, (pair.p2.x-pair.p1.x) > 0
				switch {
				case p2BigY && p2BigX:
					grid[pair.p1.x+i][pair.p1.y+i] += 1
				case p2BigY && !p2BigX:
					grid[pair.p1.x-i][pair.p1.y+i] += 1
				case !p2BigY && p2BigX:
					grid[pair.p1.x+i][pair.p1.y-i] += 1
				case !p2BigY && !p2BigX:
					grid[pair.p1.x-i][pair.p1.y-i] += 1
				}

			}
		}
	}

	return overlap(grid)
}

func overlap(grid [1000][1000]int) int {
	total := 0
	for _, row := range grid {
		for _, val := range row {
			if val > 1 {
				total += 1
			}
		}
	}

	return total
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func readFile(path string) []Line {
	file, err := os.Open(path)
	lines := []Line{}

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		pointsString := strings.Split(scanner.Text(), " -> ")
		pointString1 := strings.Split(pointsString[0], ",")
		pointString2 := strings.Split(pointsString[1], ",")

		p1X, _ := strconv.Atoi(pointString1[0])
		p1Y, _ := strconv.Atoi(pointString1[1])
		p2X, _ := strconv.Atoi(pointString2[0])
		p2Y, _ := strconv.Atoi(pointString2[1])

		lines = append(lines, Line{Point{p1X, p1Y}, Point{p2X, p2Y}})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
