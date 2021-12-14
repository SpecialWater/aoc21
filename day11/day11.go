package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Octo struct {
	Row int
	Col int
}

type OctoState struct {
	Flashed bool
	Value   int
}

var Coordinates = make(map[Octo]OctoState)
var Flashes = 0

func Day11() {
	readFile("day11/input.txt")
	// solve1()

	fmt.Printf("Answer to Day11, Q1: %v\n", Flashes)

	ans := solve2()
	fmt.Printf("Answer to Day11, Q2: %v\n", ans)

	//scores := completeLine(incompleteLines)
	//

}

// Increase by 1 and set all flash values to false
func increaseValue() {
	for k, v := range Coordinates {
		Coordinates[k] = OctoState{false, v.Value + 1}
	}
}

func Flash() {
	for k, v := range Coordinates {
		if v.Value >= 10 && !v.Flashed {
			Coordinates[k] = OctoState{true, 0}
			Flashes += 1
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if key, ok := Coordinates[Octo{k.Row + i, k.Col - j}]; ok && !key.Flashed {
						Coordinates[Octo{k.Row + i, k.Col - j}] = OctoState{key.Flashed, key.Value + 1}
						Flash()
					}
				}
			}
		}
	}
}

func solve1() {
	for i := 0; i < 100; i++ {
		increaseValue() // First Move
		Flash()
	}
}

func solve2() int {
	counter := 1
	for {
		tmpFlashes := Flashes
		increaseValue() // First Move
		Flash()
		if Flashes-tmpFlashes == len(Coordinates) {
			return counter
		}
		counter += 1
	}
}

func readFile(path string) {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	row := 0
	for scanner.Scan() {
		for i, v := range scanner.Text() {
			val, _ := strconv.Atoi(string(v))
			Coordinates[Octo{row, i}] = OctoState{false, val}
		}
		row += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
