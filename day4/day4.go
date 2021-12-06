package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// New approach to part 2:
// Make a map, key = binary, value = bool
// while loop: count most common, remove least common, break when one key left

// Starting from 0 index

type Row struct {
	vals  []int
	bools []bool
}

type Board struct {
	Rows   []Row
	Winner bool
}

func Day4() {
	numbers, boards := readFile("day4/input.txt")
	// fmt.Println(numbers)

	part1(numbers, boards)

}

func part1(numbers []int, b []Board) {
	winner := false
	winningBoard := []Row{}

	for _, num := range numbers {
		fmt.Println(num)
		for i_board, board := range b {
			for i_rows, r := range board.Rows {
				for i_vals, v := range r.vals {
					// fmt.Println(v, num)
					if v == num {

						b[i_board].Rows[i_rows].bools[i_vals] = true
					}
				}

			}
		}
		b, winningBoard, winner = winChecker(&b)
		if winner {
			calcWinnner(winningBoard)
			// fmt.Println(unmarkedSum * num)
		}

		lastWinningBoard, lastWinningBoardFound := calcLastWInner(b)
		if lastWinningBoardFound {
			lastWinnerSum := calcWinnner(lastWinningBoard)
			fmt.Println(lastWinningBoard)
			fmt.Println(lastWinnerSum * num)
			break
		}
	}

	// fmt.Println(b)
}

func calcLastWInner(b []Board) ([]Row, bool) {
	remainingLosers := 0
	for _, board := range b {
		if !board.Winner {
			remainingLosers += 1
		}
	}

	fmt.Println("Remaining losers: ", remainingLosers)

	if remainingLosers == 1 {
		for _, board := range b {
			if !board.Winner {
				return board.Rows, true
			}
		}
	}

	return []Row{}, false
}

func winChecker(b *[]Board) ([]Board, []Row, bool) {

	for i_board, board := range *b {
		for _, r := range board.Rows {
			winner := true
			for i_vals := range r.vals {
				if !r.bools[i_vals] {
					winner = false
					break
				}
			}
			if winner {
				// winningBoard := (*b)[i_board].Rows
				(*b)[i_board].Winner = true
				// return *b, winningBoard, true
			}
		}
	}

	return *b, []Row{}, false
}

func calcWinnner(winningRows []Row) int {
	unmarkedSum := 0
	for rowNum, row := range winningRows {
		if rowNum == 5 {
			break
		}
		for i, v := range row.vals {
			if !row.bools[i] {
				unmarkedSum += v
			}
		}
	}

	return unmarkedSum
}

func readFile(path string) ([]int, []Board) {
	file, err := os.Open(path)
	numbers := []int{}
	firstRow, secondRow := true, true
	board := []Board{}
	rows := []Row{}

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Get winning numbers
		if firstRow {
			firstRow = false
			strings := strings.Split(scanner.Text(), ",")
			for _, v := range strings {
				s, _ := strconv.Atoi(v)
				numbers = append(numbers, s)
			}
		} else {
			if secondRow {
				secondRow = false
				continue
			}

			if scanner.Text() == "" {
				colsToRows(&rows)
				board = append(board, Board{rows, false})
				rows = []Row{}
				continue
			} else {
				nums := strings.Split(scanner.Text(), " ")
				row := []int{}
				rowBool := []bool{}

				for _, num := range nums {
					n, err := strconv.Atoi(num)
					if err != nil {
						continue
					} else {
						row = append(row, n)
						rowBool = append(rowBool, false)
					}
				}
				rows = append(rows, Row{row, rowBool})
			}
		}
	}

	return numbers, board
}

func colsToRows(rows *[]Row) {
	newRows := []Row{}
	for col := 0; col != 5; col++ {
		row := []int{}
		rowBool := []bool{}
		for _, r := range *rows {
			for i, v := range r.vals {
				if i == col {
					row = append(row, v)
					rowBool = append(rowBool, false)
				}
			}
		}
		newRows = append(newRows, Row{row, rowBool})
	}

	*rows = append(*rows, newRows...)
}
