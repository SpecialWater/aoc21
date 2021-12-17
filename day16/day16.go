package day16

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Packet struct {
	PacketNumber int
	Version      int
	TypeID       int
	Value        int
	SubPackets   []int
}

func Day16() {
	binary := readFile("day16/input.txt")
	fmt.Println(binary)

	//fmt.Printf("Answer to Day12, Q1: %v\n", ans1)
	//fmt.Printf("Answer to Day12, Q2: %v\n", ans2)
}

func readFile(path string) []int {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	binary := []int{}
	for scanner.Scan() {
		text := scanner.Text()

		for _, v := range text {
			i, _ := strconv.ParseUint(string(v), 16, 4)
			bin := fmt.Sprintf("%04b", i)
			for _, b := range bin {
				val, _ := strconv.Atoi(string(b))
				binary = append(binary, val)
			}

		}
	}

	return binary

}
