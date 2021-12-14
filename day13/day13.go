package day13

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type fold struct {
	direction string
	value     int
}

var instructions = make([][]bool, 1311)
var folds = []fold{}

func Day13() {
	readFile("day13/input.txt")

	solve2()
	//foldPaper(folds[0])
	//fmt.Println(len(instructions))
	//fmt.Printf("Answer to Day12, Q1: %v\n", ans1)
	//fmt.Printf("Answer to Day12, Q2: %v\n", ans2)

}

func solve2() {
	for _, v := range folds {
		foldPaper(v)
	}

	p2 := make([][]string, len(instructions))
	for i := range p2 {
		p2[i] = make([]string, len(instructions[0]))
	}

	for i, row := range p2 {
		for j := range row {
			if instructions[i][j] {
				p2[i][j] = "#"
			} else {
				p2[i][j] = "."
			}
		}
	}

	for _, row := range p2 {
		fmt.Println(row)
	}

}

func foldPaper(f fold) {

	if f.direction == "y" {
		for i := range instructions {
			for j := 0; j < f.value; j++ {
				instructions[i][j] = instructions[i][j] || instructions[i][f.value*2-j]
			}
		}

		tmp := make([][]bool, len(instructions))
		for i := range tmp {
			tmp[i] = make([]bool, f.value)
		}

		for i, row := range tmp {
			for j := range row {
				tmp[i][j] = instructions[i][j]
			}
		}
		instructions = tmp
	}

	if f.direction == "x" {
		for i := 0; i < f.value; i++ {
			for j := range instructions[i] {
				instructions[i][j] = instructions[i][j] || instructions[f.value*2-i][j]
			}
		}

		tmp := make([][]bool, f.value)
		for i := range tmp {
			tmp[i] = make([]bool, len(instructions[0]))
		}

		for i, row := range tmp {
			for j := range row {
				tmp[i][j] = instructions[i][j]
			}
		}
		instructions = tmp
	}

}

func readFile(path string) {
	file, err := os.Open(path)
	for i := range instructions {
		instructions[i] = make([]bool, 895)
	}

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	lineSkip := false

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		if len(data) == 1 && !lineSkip {
			lineSkip = true
			continue
		}

		if !lineSkip {
			x, _ := strconv.Atoi(data[0])
			y, _ := strconv.Atoi(data[1])
			instructions[x][y] = true
		} else {
			f := strings.Split(data[0], " ")
			f = strings.Split(f[2], "=")

			f_val, _ := strconv.Atoi(f[1])
			folds = append(folds, fold{f[0], f_val})
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
