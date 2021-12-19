package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type board struct {
	squares [5][5]int
	marked  [5][5]bool
	//Values in array are instantiated to default value of type; for boolean this is false
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	numbers, boards := readInput()
	for i := 0; i < len(numbers); i++ {
		currentNb := numbers[i]
		for j := 0; j < len(boards); j++ {
			boards[j] = markBoard(boards[j], currentNb)
		}
		for j := 0; j < len(boards); j++ {
			if checkVictory(boards[j]) == true {
				score := currentNb * getSumOfUnmarked(boards[j])
				return score
			}
		}
	}
	return -1
}

func part2() int {
	numbers, boards := readInput()
	for i := 0; i < len(numbers); i++ {
		currentNb := numbers[i]
		for j := 0; j < len(boards); j++ {
			boards[j] = markBoard(boards[j], currentNb)
		}
		if len(boards) == 1 {
			if checkVictory(boards[0]) == true {
				score := currentNb * getSumOfUnmarked(boards[0])
				return score
			}
		}
		var newBoards []board
		for j := 0; j < len(boards); j++ {
			if checkVictory(boards[j]) == false {
				newBoards = append(newBoards, boards[j])
			}
		}
		boards = newBoards
	}
	return -1
}

func readInput() ([]int, []board) {
	f, err := os.Open("input-04.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//Retrieve the called numbers
	var numbers []int
	split := strings.Split(lines[0], ",")
	for i := 0; i < len(split); i++ {
		number, err := strconv.Atoi(split[i])
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	//Boards occur at a regular pattern, the first line of a board is always situated at line 3 + 6n (2 + 6n in the array).

	var boards []board

	for i := 2; i < len(lines); i = i + 6 {
		b := board{}
		rowNb := 0
		for j := i; j < i+5; j++ {
			//First replace icky double spaces
			line := strings.Replace(lines[j], "  ", " ", -1)
			split2 := strings.Split(line, " ")
			//Clean up potential white space at start of line
			if split2[0] == "" {
				split2 = split2[1:]
			}
			for k := 0; k < 5; k++ {
				nb, err := strconv.Atoi(split2[k])
				if err != nil {
					log.Fatal(err)
				}
				b.squares[rowNb][k] = nb
			}
			rowNb++
		}
		boards = append(boards, b)
	}

	return numbers, boards
}

func markBoard(board board, number int) board {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board.squares[i][j] == number {
				board.marked[i][j] = true
			}
		}
	}
	return board
}

func checkVictory(board board) bool {
	//We can check a row and a column at the same time; just switch the coordinates.
	for i := 0; i < 5; i++ {
		rowWins := true
		columnWins := true
		for j := 0; j < 5; j++ {
			if board.marked[i][j] == false {
				rowWins = false
			}
		}
		for j := 0; j < 5; j++ {
			if board.marked[j][i] == false {
				columnWins = false
			}
		}
		if rowWins == true {
			return true
		}
		if columnWins == true {
			return true
		}
	}
	return false
}

func getSumOfUnmarked(board board) int {
	result := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board.marked[i][j] == false {
				result = result + board.squares[i][j]
			}
		}
	}
	return result
}
