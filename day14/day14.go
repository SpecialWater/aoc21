package day14

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Day14() {
	readFile("day14/input.txt")

	//fmt.Printf("Answer to Day12, Q1: %v\n", ans1)
	//fmt.Printf("Answer to Day12, Q2: %v\n", ans2)

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
		data := strings.Split(scanner.Text(), ",")

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
