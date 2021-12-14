package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type LowPoint struct {
	Row int
	Col int
}

func Day9() {
	data := readFile("day9/input.txt")

	ans1, lowPoints := solve1(data)
	fmt.Printf("Answer to Day9, Q1: %v\n", ans1)

	fmt.Printf("Answer to Day9, Q2: %v\n", solve2(data, lowPoints))
}

func solve1(data [][]int) (int, []LowPoint) {
	sum := 0
	rowLen := len(data)
	colLen := len(data[0])
	lowPoints := []LowPoint{}

	for row, _ := range data {
		for col, _ := range data[row] {
			switch {
			case row == 0 && col == (colLen-1): // Top Right corner
				if checkLow(data[row][col], data[row+1][col], data[row][colLen-2]) {
					sum += data[row][col] + 1
					lowPoints = append(lowPoints, LowPoint{row, col})
				}
			case row == 0 && col == 0: // Top Left Corner
				if checkLow(data[row][col], data[row+1][col], data[row][col+1]) {
					sum += data[row][col] + 1
					lowPoints = append(lowPoints, LowPoint{row, col})
				}
			case row == (rowLen-1) && col == 0: // Bottom Left Corner
				if checkLow(data[row][col], data[row-1][col], data[row][col+1]) {
					sum += data[row][col] + 1
					lowPoints = append(lowPoints, LowPoint{row, col})
				}
			case row == (rowLen-1) && col == (colLen-1): // Bottom Right Corner
				if checkLow(data[row][col], data[row-1][col], data[row][colLen-2]) {
					sum += data[row][col] + 1
					lowPoints = append(lowPoints, LowPoint{row, col})
				}
			case row == 0: // Top Row
				if checkLow(data[row][col], data[row+1][col], data[row][col+1], data[row][col-1]) {
					sum += data[row][col] + 1
					lowPoints = append(lowPoints, LowPoint{row, col})
				}
			case col == 0: // Left Column
				if checkLow(data[row][col], data[row+1][col], data[row-1][col], data[row][col+1]) {
					sum += data[row][col] + 1
					lowPoints = append(lowPoints, LowPoint{row, col})
				}
			case col == (colLen - 1): // Right Column
				if checkLow(data[row][col], data[row+1][col], data[row-1][col], data[row][col-1]) {
					sum += data[row][col] + 1
					lowPoints = append(lowPoints, LowPoint{row, col})
				}
			case row == (rowLen - 1): // Bottom Row
				if checkLow(data[row][col], data[row-1][col], data[row][col-1], data[row][col+1]) {
					sum += data[row][col] + 1
					lowPoints = append(lowPoints, LowPoint{row, col})
				}
			default:
				if checkLow(data[row][col], data[row-1][col], data[row+1][col], data[row][col-1], data[row][col+1]) {
					sum += data[row][col] + 1
					lowPoints = append(lowPoints, LowPoint{row, col})
				}
			}
		}
	}
	return sum, lowPoints
}

func checkLow(val int, surround ...int) bool {
	lowPoint := true
	for _, v := range surround {
		if val >= v {
			lowPoint = false
		}
	}

	return lowPoint
}

func solve2(data [][]int, lowPoints []LowPoint) int {
	explored := make(map[LowPoint]map[LowPoint]bool)
	rowLen := len(data)
	colLen := len(data[0])

	for _, v := range lowPoints {
		explored[v] = make(map[LowPoint]bool)
		basin(rowLen, colLen, data, v, v, explored)
	}

	basinSizes := []int{}
	for _, v := range lowPoints {
		basinSizes = append(basinSizes, len(explored[v]))
	}

	sort.Ints(basinSizes)
	bLen := len(basinSizes)

	return basinSizes[bLen-1] * basinSizes[bLen-2] * basinSizes[bLen-3]
}

func basin(rowLen, colLen int, data [][]int, origPoint, point LowPoint, explored map[LowPoint]map[LowPoint]bool) {
	up, down := LowPoint{point.Row - 1, point.Col}, LowPoint{point.Row + 1, point.Col}
	left, right := LowPoint{point.Row, point.Col - 1}, LowPoint{point.Row, point.Col + 1}

	if up.Row >= 0 && data[up.Row][up.Col] != 9 && !explored[origPoint][up] {
		explored[origPoint][up] = true
		basin(rowLen, colLen, data, origPoint, up, explored)
	}

	if down.Row < rowLen && data[down.Row][down.Col] != 9 && !explored[origPoint][down] {
		explored[origPoint][down] = true
		basin(rowLen, colLen, data, origPoint, down, explored)
	}

	if left.Col >= 0 && data[left.Row][left.Col] != 9 && !explored[origPoint][left] {
		explored[origPoint][left] = true
		basin(rowLen, colLen, data, origPoint, left, explored)
	}

	if right.Col < colLen && data[right.Row][right.Col] != 9 && !explored[origPoint][right] {
		explored[origPoint][right] = true
		basin(rowLen, colLen, data, origPoint, right, explored)
	}

}

func readFile(path string) [][]int {
	file, err := os.Open(path)
	data := [][]int{}

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {

		tmp := strings.Split(scanner.Text(), "")
		tmpInt := []int{}
		for _, v := range tmp {
			i, _ := strconv.Atoi(v)
			tmpInt = append(tmpInt, i)
		}
		data = append(data, tmpInt)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}
