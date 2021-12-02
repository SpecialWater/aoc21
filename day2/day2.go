package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Commands struct {
	Command string
	Value   int
}

func Day2() {
	text, err := readFile("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	commands := buildStruct(text)

	depth, horizontal := navigate(commands)
	fmt.Printf("Answer to Day1, Q1: %v\n", depth*horizontal)

	depth2, horizontal2 := navigate2(commands)
	fmt.Printf("Answer to Day1, Q1: %v\n", depth2*horizontal2)

}

func navigate(c []Commands) (int, int) {
	h, d := 0, 0
	for _, v := range c {
		switch c := v.Command; c {
		case "down":
			d += v.Value
		case "up":
			d -= v.Value
		case "forward":
			h += v.Value
		}
	}

	return d, h
}

func navigate2(c []Commands) (int, int) {
	h, d, a := 0, 0, 0
	for _, v := range c {
		switch c := v.Command; c {
		case "down":
			a += v.Value
		case "up":
			a -= v.Value
		case "forward":
			h += v.Value
			d += a * v.Value
		}
	}

	return d, h
}

func buildStruct(c []string) []Commands {
	var commands = []Commands{}
	for _, v := range c {
		split := strings.Split(v, " ")
		value, _ := strconv.Atoi(split[1])
		tmp := Commands{split[0], value}
		commands = append(commands, tmp)
	}

	return commands
}

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	contents := []string{}

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return contents, nil
}
