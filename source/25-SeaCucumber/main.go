package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type point struct {
	row    int
	column int
}

func main() {
	solve()
}

func solve() {
	east, south, height, width := readInput("input-25.txt")
	done := false
	steps := 0
	for !done {
		movements := 0
		newEast := make(map[point]bool)
		newSouth := make(map[point]bool)
		for p, _ := range east {
			newP := point{}
			newP.row = p.row
			if p.column == width-1 {
				newP.column = 0
			} else {
				newP.column = p.column + 1
			}
			_, inEast := east[newP]
			_, inSouth := south[newP]
			if inEast || inSouth {
				newEast[p] = true
			} else {
				newEast[newP] = true
				movements = movements + 1
			}
		}
		east = newEast
		for p, _ := range south {
			newP := point{}
			if p.row == height-1 {
				newP.row = 0
			} else {
				newP.row = p.row + 1
			}
			newP.column = p.column
			_, inEast := east[newP]
			_, inSouth := south[newP]
			if inEast || inSouth {
				newSouth[p] = true
			} else {
				newSouth[newP] = true
				movements = movements + 1
			}
		}
		south = newSouth
		steps = steps + 1
		if movements == 0 || steps >= 10000 {
			done = true
		}
	}
	fmt.Printf("Result: %d\n", steps)
}

func readInput(path string) (map[point]bool, map[point]bool, int, int) {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	east := make(map[point]bool)
	south := make(map[point]bool)

	row := 0

	width := 0
	height := 0

	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), "")
		width = len(splitLine)
		for i := 0; i < len(splitLine); i++ {
			if splitLine[i] == ">" {
				p := point{}
				p.row = row
				p.column = i
				east[p] = true
			} else if splitLine[i] == "v" {
				p := point{}
				p.row = row
				p.column = i
				south[p] = true
			}
		}
		row = row + 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	height = row

	return east, south, height, width
}
