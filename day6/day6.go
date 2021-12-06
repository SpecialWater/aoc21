package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Fish struct {
	Cycle int
}

type Colony []Fish

func simulateDays(fish Colony, days int) int {
	population := [9]int{}
	for _, v := range fish {
		population[v.Cycle] += 1
	}

	for i := 0; i < days; i++ {
		populationPriorDay := population
		population[0] = populationPriorDay[1]
		population[1] = populationPriorDay[2]
		population[2] = populationPriorDay[3]
		population[3] = populationPriorDay[4]
		population[4] = populationPriorDay[5]
		population[5] = populationPriorDay[6]
		population[6] = populationPriorDay[7] + populationPriorDay[0]
		population[7] = populationPriorDay[8]
		population[8] = populationPriorDay[0]
	}

	total := 0
	for _, v := range population {
		total += v
	}

	return total

}

func Day6() {
	colony := readFile("day6/input.txt")
	simulateDays(colony, 256)
	fmt.Printf("Answer to Day6, Q1: %v\n", simulateDays(colony, 80))
	fmt.Printf("Answer to Day6, Q2: %v\n", simulateDays(colony, 256))
}

func readFile(path string) []Fish {
	file, err := os.Open(path)
	sliceStrings := []string{}
	fish := Colony{}

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		sliceStrings = strings.Split(scanner.Text(), ",")
	}

	for _, s := range sliceStrings {
		v, _ := strconv.Atoi(s)
		fish = append(fish, Fish{v})

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return fish
}
