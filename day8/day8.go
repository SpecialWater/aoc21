package day8

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

type Data struct {
	inputs         []string
	outputs        []string
	outputsDecoded []string
	outputsInts    []int
}

var sevenDigit = make(map[string]int)

func Day8() {
	data := readFile("day8/input.txt")
	sevenDigit["abcefg"] = 0
	sevenDigit["cf"] = 1
	sevenDigit["acdeg"] = 2
	sevenDigit["acdfg"] = 3
	sevenDigit["bcdf"] = 4
	sevenDigit["abdfg"] = 5
	sevenDigit["abdefg"] = 6
	sevenDigit["acf"] = 7
	sevenDigit["abcdefg"] = 8
	sevenDigit["abcdfg"] = 9

	data = Decrypt(data)

	fmt.Printf("Answer to Day8, Q1: %v\n", problemOne(data))
	fmt.Printf("Answer to Day8, Q2: %v\n", problemTwo(data))
}

// swap b and d

func problemOne(data []Data) int {
	total := 0
	for _, row := range data {
		for _, v := range row.outputsInts {
			if v == 1 || v == 4 || v == 7 || v == 8 {
				total += 1
			}
		}
	}

	return total
}

func problemTwo(data []Data) float64 {
	total := 0.0
	for _, row := range data {
		// fmt.Println(row.outputs, row.outputsDecoded, row.outputsInts)
		for i, v := range row.outputsInts {
			total += float64(v) * math.Pow10(3-i)
		}
	}

	return total
}

// Known numbers by length: 1, 4, 7, 8
// Known wires: a, c, e, f, g,
// Numbers with 5 length: 2, 3, 5
// Numbers with 6 lenngth: 0, 6, 9
func Decrypt(data []Data) []Data {

	for i, v := range data {
		decryptor := make(map[string]string)
		solve(v.inputs, decryptor)
		// fmt.Println(decryptor)
		data[i].outputsDecoded, data[i].outputsInts = decryptOutput(v.outputs, decryptor)
	}

	return data
}

func decryptOutput(outputCoded []string, decryptor map[string]string) ([]string, []int) {
	outputDecoded := []string{}
	outputInts := []int{}
	for _, str := range outputCoded {
		decodedString := ""
		for _, c := range str {
			for k, v := range decryptor {
				if string(c) == v {
					decodedString += k
					break
				}
			}
		}
		decodedString = SortString(decodedString)
		outputDecoded = append(outputDecoded, decodedString)
		outputInts = append(outputInts, sevenDigit[decodedString])
	}

	return outputDecoded, outputInts
}

func solve(input []string, decryptor map[string]string) {
	solveCF(input, decryptor)
	solveA(input, decryptor)
	solveBD(input, decryptor)
	solveG(input, decryptor)
	solveE(input, decryptor)
	solveDC(input, decryptor)
	solveBandD(input, decryptor)
}

func solveCF(input []string, decryptor map[string]string) {
	for _, v := range input {
		if len(v) == 2 {
			decryptor["c"] = v
			decryptor["f"] = v
		}
	}
}

func solveBD(input []string, decryptor map[string]string) {
	for _, v := range input {
		if len(v) == 4 {
			for _, c := range v {
				if !strings.Contains(decryptor["c"], string(c)) {
					decryptor["b"] = decryptor["b"] + string(c)
					decryptor["d"] = decryptor["d"] + string(c)
				}
			}
		}
	}
}

func solveA(input []string, decryptor map[string]string) {
	for _, v := range input {
		if len(v) == 3 {
			for _, c := range v {
				if !strings.Contains(decryptor["c"], string(c)) {
					decryptor["a"] = string(c)
				}
			}
		}
	}
}

func solveG(input []string, decryptor map[string]string) {
	for _, v := range input {
		if len(v) == 6 {
			unknown := 0
			for _, c := range v {
				knownVals := decryptor["a"] + decryptor["c"] + decryptor["b"]
				if !strings.Contains(knownVals, string(c)) {
					unknown += 1
				}
			}
			if unknown == 1 {
				for _, c := range v {
					knownVals := decryptor["a"] + decryptor["c"] + decryptor["b"]
					if !strings.Contains(knownVals, string(c)) {
						decryptor["g"] = string(c)
					}
				}
			}
		}
	}
}

func solveE(input []string, decryptor map[string]string) {
	str := "abcdefg"
	known := ""
	for _, v := range decryptor {
		known += v
	}

	for _, s := range str {
		if !strings.Contains(known, string(s)) {
			decryptor["e"] = string(s)
		}
	}
}

func solveDC(input []string, decryptor map[string]string) {
	// knowns := "aeg"
	knownsCoded := decryptor["a"] + decryptor["e"] + decryptor["g"]
	for _, v := range input {
		if len(v) == 5 {
			unknown := 0
			for _, c := range v {
				if !strings.Contains(knownsCoded, string(c)) {
					unknown += 1
				}
			}
			// On number Two - c and d unknown
			if unknown == 2 {
				for _, c := range v {
					if !strings.Contains(knownsCoded, string(c)) {
						if strings.ContainsAny(decryptor["c"], string(c)) {
							tmp := decryptor["c"]
							decryptor["c"] = string(c)
							for _, x := range tmp {
								if string(x) != string(c) {
									decryptor["f"] = string(x)
								}
							}
							break
						}
					}
				}
			}
		}
	}
}

func solveBandD(input []string, decryptor map[string]string) {
	knownsCoded := decryptor["a"] + decryptor["c"] + decryptor["e"] + decryptor["f"] + decryptor["g"]
	for _, v := range input {
		if len(v) == 6 {
			unknown := 0
			for _, c := range v {
				if !strings.Contains(knownsCoded, string(c)) {
					unknown += 1
				}
			}
			// On number Two - b and d unknown
			if unknown == 1 {
				for _, c := range v {
					if !strings.Contains(knownsCoded, string(c)) {
						tmp := decryptor["b"]
						decryptor["b"] = string(c)
						for _, x := range tmp {
							if string(x) != string(c) {
								decryptor["d"] = string(x)
							}
						}
						break
					}
				}
			}
		}
	}
}

func readFile(path string) []Data {
	file, err := os.Open(path)
	data := []Data{}

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), " | ")
		inputs := strings.Split(tmp[0], " ")
		outputs := strings.Split(tmp[1], " ")
		data = append(data, Data{inputs, outputs, []string{}, []int{}})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
