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
	part1()
}

func part1() {
	numbers, boards := readInput()
	fmt.Println(numbers)
	fmt.Println(boards)
}

func readInput() ([]int, []board) {
	f, err := os.Open("input-04-example.txt")

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
