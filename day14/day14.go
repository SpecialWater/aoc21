package day14

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var polyBuild = make(map[string]string)

func Day14() {
	polyCode := readFile("day14/input.txt")

	fmt.Printf("Answer to Day14, Q1: %v\n", solve(polyCode, 10))
	fmt.Printf("Answer to Day14, Q2: %v\n", solve(polyCode, 40))

}

func solve(polyCode string, iterations int) int {
	itSmall := iterations / 2
	polyCounts := make(map[string]map[string]int)

	// iterate halfway through total iterations
	// Get Strings generated from 2 characters
	for k := range polyBuild {
		code := k
		for i := 0; i < itSmall; i++ {
			code = iterate(code)
		}
		polyCounts[k] = getCount(code)
	}

	smallPolyCode := polyCode
	for i := 0; i < itSmall; i++ {
		smallPolyCode = iterate(smallPolyCode)
	}

	totals := countTotal(smallPolyCode, polyCounts)

	return getAnswer(totals)
}

func getAnswer(totalsMap map[string]int) int {
	min := 999999999999999999
	max := 0

	for _, v := range totalsMap {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return max - min

}

func countTotal(str string, polyCounts map[string]map[string]int) map[string]int {
	countTotal := make(map[string]int)

	for i := 0; i < len(str)-1; i++ {
		s := string(str[i]) + string(str[i+1])
		counts := polyCounts[s]
		for k, v := range counts {
			countTotal[k] += v
		}
	}

	for i := 0; i < len(str)-1; i++ {
		countTotal[string(str[i])] += 1
	}

	countTotal[string(str[len(str)-1])] += 1

	return countTotal

}

func getCount(str string) map[string]int {
	charCounts := make(map[string]int)
	for _, s := range str {
		charCounts[string(s)] += 1
	}

	charCounts[string(str[0])] -= 1
	charCounts[string(str[len(str)-1])] -= 1

	return charCounts
}

func iterate(polyCode string) string {
	var sb strings.Builder
	for i := 0; i < len(polyCode)-1; i++ {
		s := polyBuild[string(polyCode[i])+string(polyCode[i+1])]
		sb.WriteString(s)
	}
	sb.WriteString(string(polyCode[len(polyCode)-1]))

	return sb.String()
}

func readFile(path string) string {
	file, err := os.Open(path)
	polyCode := ""

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	lineSkip := false

	for scanner.Scan() {
		data := scanner.Text()
		if len(data) == 0 && !lineSkip {
			lineSkip = true
			continue
		}

		if !lineSkip {
			polyCode = data
		} else {

			inst := strings.Split(data, " -> ")
			firstChar := inst[0][0]
			polyBuild[inst[0]] = string(firstChar) + inst[1]
		}

	}

	return polyCode

}
