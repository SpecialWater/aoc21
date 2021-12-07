package day7

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day7() {
	positions := readFile("day7/input.txt")
	fuelSolve(positions)
	fuelSolve2(positions)

	fmt.Printf("Answer to Day7, Q1: %v\n", fuelSolve(positions))
	fmt.Printf("Answer to Day7, Q2: %v\n", fuelSolve2(positions))
}

func median(s []int) int {
	sLen := len(s)
	if sLen%2 == 0 {
		return (s[sLen/2] + s[sLen/2-1]) / 2
	} else {
		return s[(sLen-1)/2]
	}
}

func average(s []int) int {
	avg := 0
	for _, v := range s {
		avg += v
	}
	return avg / len(s)
}

func fuelSolve2(s []int) int {
	bestFuel := 99999999999

	for i := 0; i <= 2000; i++ {
		fuel := 0
		for _, v := range s {
			l := int(math.Abs(float64(i - v)))
			fuel += consecIntSum(l, 1, l)
		}
		if fuel < bestFuel {
			//fmt.Printf("Iteration: %v, Fuel: %v\n", i, fuel)
			bestFuel = fuel
		}
	}
	return bestFuel
}

func consecIntSum(l, first, last int) int {
	res := int(float64(l) * float64(first+last) / float64(2))
	return res
}

func fuelSolve(s []int) int {
	m := median(s)
	bestFuel := 99999999

	for i := -1; i <= 1; i++ {
		fuel := 0
		for _, v := range s {
			fuel += int(math.Abs(float64(v - m + i)))
		}
		if fuel < bestFuel {
			//fmt.Printf("Iteration: %v, Fuel: %v\n", i, fuel)
			bestFuel = fuel
		}
	}

	return bestFuel
}

func readFile(path string) []int {
	file, err := os.Open(path)
	positions := []int{}
	tmp := []string{}

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		tmp = strings.Split(scanner.Text(), ",")
	}

	for _, s := range tmp {
		v, _ := strconv.Atoi(s)
		positions = append(positions, v)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(positions)

	return positions
}
