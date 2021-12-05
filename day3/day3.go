package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// New approach to part 2:
// Make a map, key = binary, value = bool
// while loop: count most common, remove least common, break when one key left

// Starting from 0 index
var BitLength = 11

func Day3() {
	text, oxyCo2, err := readFile("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	e, g := gammaEpsilon(sumBitColumn(text, BitLength))
	fmt.Printf("Answer to Day3, Q1: %v\n", e*g)

	oxy, co2 := OxyCo2(oxyCo2)
	fmt.Printf("Answer to Day3, Q2: %v\n", oxy*co2)

}

func sumBitColumn(bits string, bitLength int) (map[int]int, int) {
	columnTotals := make(map[int]int)

	column, rowCount := 0, 0
	for _, v := range bits {
		val, _ := strconv.Atoi(string(v))
		columnTotals[column] += val

		if column == bitLength {
			column = 0
			rowCount += 1
		} else {
			column += 1
		}
	}

	fmt.Println(columnTotals, rowCount)

	return columnTotals, rowCount
}

func gammaEpsilon(columnTotals map[int]int, rowCount int) (int, int) {
	epsilon, gamma := make([]byte, BitLength+1), make([]byte, BitLength+1)
	for k, v := range columnTotals {
		if v > (rowCount / 2) {
			epsilon[k] = '1'
			gamma[k] = '0'
		} else {
			epsilon[k] = '0'
			gamma[k] += '1'
		}
	}

	e, _ := strconv.ParseInt(string(epsilon), 2, 64)
	g, _ := strconv.ParseInt(string(gamma), 2, 64)

	return int(e), int(g)
}

func OxyCo2(values map[string]bool) (int, int) {
	valuesOxy := make(map[string]bool)
	valuesCo2 := make(map[string]bool)

	for key, value := range values {
		valuesOxy[key] = value
		valuesCo2[key] = value
	}

	day2Filter(valuesOxy, "1", "0")
	day2Filter(valuesCo2, "0", "1")

	ret1, ret2 := "", ""
	for k := range valuesOxy {
		ret1 = k
	}

	for k := range valuesCo2 {
		ret2 = k
	}

	o, _ := strconv.ParseInt(string(ret1), 2, 64)
	c, _ := strconv.ParseInt(string(ret2), 2, 64)

	return int(o), int(c)
}

func day2Filter(vals map[string]bool, keep, del string) {
	column := 0
	for len(vals) != 1 {
		total := 0
		for k := range vals {
			if string(k[column]) == "1" {
				total += 1
			}
		}

		bigSmall := float64(total) >= float64(len(vals))/float64(2)
		for k := range vals {
			if v := string(k[column]); bigSmall {
				if v == del {
					delete(vals, k)
				}
			} else {
				if v == keep {
					delete(vals, k)
				}
			}
		}
		column += 1
	}
}

func readFile(path string) (string, map[string]bool, error) {
	file, err := os.Open(path)
	oxyCo2 := make(map[string]bool)
	var contents string

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		contents += scanner.Text()
		oxyCo2[scanner.Text()] = true
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return contents, oxyCo2, nil
}
