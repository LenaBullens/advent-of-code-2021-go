package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part1(readInput("input-11.txt"), 100))
	fmt.Println(part2(readInput("input-11.txt")))
}

func part1(octos [][]int, steps int) int {
	var totalFlashes int
	for i := 0; i < steps; i++ {
		var extraFlashes int
		octos, extraFlashes = step(octos)
		totalFlashes = totalFlashes + extraFlashes

		//Print octos
		fmt.Println("---Step " + strconv.Itoa(i+1) + "---")
		for j := 0; j < len(octos); j++ {
			for k := 0; k < len(octos[0]); k++ {
				fmt.Print(octos[j][k])
			}
			fmt.Println("")
		}
		fmt.Println("")
	}

	return totalFlashes
}

func part2(octos [][]int) int {
	done := false
	step := 1

	for !done {
		octos, done = step2(octos)

		//Print octos
		fmt.Println("---Step " + strconv.Itoa(step) + "---")
		for j := 0; j < len(octos); j++ {
			for k := 0; k < len(octos[0]); k++ {
				fmt.Print(octos[j][k])
			}
			fmt.Println("")
		}
		fmt.Println("")

		if done {
			return step
		}
		step = step + 1
	}

	return -1
}

func step(octos [][]int) ([][]int, int) {
	var shinies [10][10]bool
	var nbOfFlashes int

	//Increase level
	for i := 0; i < len(octos); i++ {
		for j := 0; j < len(octos[0]); j++ {
			octos[i][j] = octos[i][j] + 1
		}
	}

	//Flash
	for i := 0; i < len(octos); i++ {
		for j := 0; j < len(octos[0]); j++ {
			if octos[i][j] > 9 && !shinies[i][j] {
				var extraFlashes int
				shinies, extraFlashes = flash(octos, shinies, i, j)
				nbOfFlashes = nbOfFlashes + extraFlashes
			}
		}
	}

	//Reset shinies
	for i := 0; i < len(shinies); i++ {
		for j := 0; j < len(shinies[i]); j++ {
			if shinies[i][j] {
				octos[i][j] = 0
			}
		}
	}

	return octos, nbOfFlashes
}

func step2(octos [][]int) ([][]int, bool) {
	var shinies [10][10]bool
	done := false

	//Increase level
	for i := 0; i < len(octos); i++ {
		for j := 0; j < len(octos[0]); j++ {
			octos[i][j] = octos[i][j] + 1
		}
	}

	//Flash
	for i := 0; i < len(octos); i++ {
		for j := 0; j < len(octos[0]); j++ {
			if octos[i][j] > 9 && !shinies[i][j] {
				shinies, _ = flash(octos, shinies, i, j)
			}
		}
	}

	var amountOfShinies int
	//Check how many flashed and reset shinies
	for i := 0; i < len(shinies); i++ {
		for j := 0; j < len(shinies[i]); j++ {
			if shinies[i][j] {
				octos[i][j] = 0
				amountOfShinies = amountOfShinies + 1
			}
		}
	}

	if amountOfShinies >= 100 {
		done = true
	}

	return octos, done
}

func flash(octos [][]int, shinies [10][10]bool, i int, j int) ([10][10]bool, int) {
	//Flash of square itself.
	nbOfFlashes := 1
	shinies[i][j] = true

	//Check surrounding squares
	if i > 0 {
		octos[i-1][j] = octos[i-1][j] + 1
		if !shinies[i-1][j] && octos[i-1][j] > 9 {
			var extraFlashes int
			shinies, extraFlashes = flash(octos, shinies, i-1, j)
			nbOfFlashes = nbOfFlashes + extraFlashes
		}
	}
	if i < len(octos)-1 {
		octos[i+1][j] = octos[i+1][j] + 1
		if !shinies[i+1][j] && octos[i+1][j] > 9 {
			var extraFlashes int
			shinies, extraFlashes = flash(octos, shinies, i+1, j)
			nbOfFlashes = nbOfFlashes + extraFlashes
		}
	}
	if j > 0 {
		octos[i][j-1] = octos[i][j-1] + 1
		if !shinies[i][j-1] && octos[i][j-1] > 9 {
			var extraFlashes int
			shinies, extraFlashes = flash(octos, shinies, i, j-1)
			nbOfFlashes = nbOfFlashes + extraFlashes
		}
	}
	if j < len(octos[0])-1 {
		octos[i][j+1] = octos[i][j+1] + 1
		if !shinies[i][j+1] && octos[i][j+1] > 9 {
			var extraFlashes int
			shinies, extraFlashes = flash(octos, shinies, i, j+1)
			nbOfFlashes = nbOfFlashes + extraFlashes
		}
	}
	if i > 0 && j > 0 {
		octos[i-1][j-1] = octos[i-1][j-1] + 1
		if !shinies[i-1][j-1] && octos[i-1][j-1] > 9 {
			var extraFlashes int
			shinies, extraFlashes = flash(octos, shinies, i-1, j-1)
			nbOfFlashes = nbOfFlashes + extraFlashes
		}
	}
	if i < len(octos)-1 && j > 0 {
		octos[i+1][j-1] = octos[i+1][j-1] + 1
		if !shinies[i+1][j-1] && octos[i+1][j-1] > 9 {
			var extraFlashes int
			shinies, extraFlashes = flash(octos, shinies, i+1, j-1)
			nbOfFlashes = nbOfFlashes + extraFlashes
		}
	}
	if i > 0 && j < len(octos[0])-1 {
		octos[i-1][j+1] = octos[i-1][j+1] + 1
		if !shinies[i-1][j+1] && octos[i-1][j+1] > 9 {
			var extraFlashes int
			shinies, extraFlashes = flash(octos, shinies, i-1, j+1)
			nbOfFlashes = nbOfFlashes + extraFlashes
		}
	}
	if i < len(octos)-1 && j < len(octos[0])-1 {
		octos[i+1][j+1] = octos[i+1][j+1] + 1
		if !shinies[i+1][j+1] && octos[i+1][j+1] > 9 {
			var extraFlashes int
			shinies, extraFlashes = flash(octos, shinies, i+1, j+1)
			nbOfFlashes = nbOfFlashes + extraFlashes
		}
	}
	return shinies, nbOfFlashes
}

func readInput(path string) [][]int {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var octos [][]int

	for scanner.Scan() {
		stringLine := strings.Split(scanner.Text(), "")
		var octoLine []int
		for i := 0; i < len(stringLine); i++ {
			value, err := strconv.Atoi(stringLine[i])
			if err != nil {
				log.Fatal(err)
			}
			octoLine = append(octoLine, value)
		}
		octos = append(octos, octoLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return octos
}
