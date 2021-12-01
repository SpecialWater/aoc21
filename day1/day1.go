package day1

import (
	"fmt"
	"log"
	"os"
)

func Day1() {
	data, err := readFile("day1/input.txt")
	if err != nil {
		log.Println("Error reading .txt file", err)
	}

	answer1 := solver1(data)
	fmt.Printf("Answer to Day1, Q1: %v\n", answer1)

	answer2 := solver2(data)
	fmt.Printf("Answer to Day1, Q2: %v", answer2)
}

func solver1(data []int) int {
	increase := 0

	for i, v := range data {
		if size := len(data); (i + 1) == size {
			break
		}
		if data[i+1] > v {
			increase += 1
		}
	}

	return increase
}

func solver2(data []int) int {
	increase := 0

	for i := range data {
		if size := len(data); (i + 3) == size {
			break
		}
		shared := data[i+1] + data[i+2]
		avgCurrent := data[i] + shared
		avgNext := shared + data[i+3]

		if avgCurrent < avgNext {
			increase += 1
		}
	}

	return increase
}

func readFile(file string) ([]int, error) {
	var data []int

	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	for {
		var i int
		var n int
		n, err = fmt.Fscanln(f, &i)
		if n == 0 || err != nil {
			break
		}
		data = append(data, i)
	}

	return data, nil
}
