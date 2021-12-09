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
	fmt.Println(part1(readInput("input-09.txt")))
}

func part1(grid [][]int) int {
	height := len(grid)
	width := len(grid[0])

	var riskLevel int

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			riskLevel = riskLevel + calculateRiskLevel(grid, i, j)
		}
	}
	return riskLevel
}

func calculateRiskLevel(grid [][]int, i int, j int) int {
	if isLowPoint(grid, i, j) {
		return 1 + grid[i][j]
	}
	return 0
}

func isLowPoint(grid [][]int, i int, j int) bool {
	if grid[i][j] == 9 {
		return false
	}
	if i > 0 && grid[i-1][j] <= grid[i][j] {
		return false
	}
	if i < len(grid)-1 && grid[i+1][j] <= grid[i][j] {
		return false
	}
	if j > 0 && grid[i][j-1] <= grid[i][j] {
		return false
	}
	if j < len(grid[0])-1 && grid[i][j+1] <= grid[i][j] {
		return false
	}
	return true
}

func readInput(path string) [][]int {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var grid [][]int

	for scanner.Scan() {
		lineStrings := strings.Split(scanner.Text(), "")
		var line []int
		for i := 0; i < len(lineStrings); i++ {
			value, err := strconv.Atoi(lineStrings[i])
			if err != nil {
				log.Fatal(err)
			}
			line = append(line, value)
		}
		grid = append(grid, line)
	}

	return grid
}
