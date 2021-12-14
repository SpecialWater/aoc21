package day10

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
)

var PointMap = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var PointMap2 = map[string]int{
	"(": 1,
	"[": 2,
	"{": 3,
	"<": 4,
}

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	// FIXME: What do we do if the stack is empty, though?

	l := len(s)
	return s[:l-1], s[l-1]
}

func Day10() {
	data := readFile("day10/input.txt")

	ans, incompleteLines := solve1(data)
	fmt.Printf("Answer to Day10, Q1: %v\n", ans)

	scores := completeLine(incompleteLines)
	fmt.Printf("Answer to Day10, Q2: %v\n", scores)

}

func solve1(data []string) (int, []stack) {
	var err error
	incomplete := []stack{}
	score := 0

	for _, s := range data {
		line := stack{}
		lineOK := true
		for _, char := range s {
			c := string(char)
			line, err = checkError(line, c)
			if err != nil {
				score += PointMap[c]
				lineOK = false
				break
			}
		}
		if lineOK {
			incomplete = append(incomplete, line)
		}
	}
	return score, incomplete
}

func checkError(line stack, c string) (stack, error) {
	if len(line) == 0 {
		line = line.Push(c)
	} else if c == "<" || c == "{" || c == "[" || c == "(" {
		line = line.Push(c)
	} else {
		switch c {
		case ">":
			if line[len(line)-1] == "<" {
				line, _ = line.Pop()
			} else {
				return line, errors.New("Invalid Syntax")
			}
		case "}":
			if line[len(line)-1] == "{" {
				line, _ = line.Pop()
			} else {
				return line, errors.New("Invalid Syntax")
			}
		case ")":
			if line[len(line)-1] == "(" {
				line, _ = line.Pop()
			} else {
				return line, errors.New("Invalid Syntax")
			}
		case "]":
			if line[len(line)-1] == "[" {
				line, _ = line.Pop()
			} else {
				return line, errors.New("Invalid Syntax")
			}
		}
	}

	return line, nil
}

func completeLine(line []stack) int {

	scores := []int{}
	for _, str := range line {
		score := 0
		for i := len(str) - 1; i >= 0; i-- {
			score = score*5 + PointMap2[str[i]]
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
}

func readFile(path string) []string {
	file, err := os.Open(path)
	data := []string{}

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}
