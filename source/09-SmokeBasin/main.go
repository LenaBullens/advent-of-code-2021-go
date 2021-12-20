package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	i int
	j int
}

func main() {
	fmt.Println(part1(readInput("input-09.txt")))
	fmt.Println(part2(readInput("input-09.txt")))
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

func part2(grid [][]int) int {
	//First identify the low-points
	var lowPoints []point

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if isLowPoint(grid, i, j) {
				point := point{i, j}
				lowPoints = append(lowPoints, point)
			}
		}
	}

	var sizes []int

	for k := 0; k < len(lowPoints); k++ {
		basinPoints := determineBasin(grid, lowPoints[k])
		sizes = append(sizes, len(basinPoints))
	}

	sort.Ints(sizes)
	return sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]
}

func determineBasin(grid [][]int, lowPoint point) []point {
	var basinPoints []point
	pointsToCheck := make(map[point]bool)
	pointsToCheck[lowPoint] = true
	checkedPoints := make(map[point]bool)

	for len(pointsToCheck) > 0 {
		newPointsToCheck := make(map[point]bool)

		for k := range pointsToCheck {
			higherPoints := getSurroundingHigherPoints(grid, k)
			basinPoints = append(basinPoints, k)
			checkedPoints[k] = true
			for l := 0; l < len(higherPoints); l++ {
				if checkedPoints[higherPoints[l]] == false {
					newPointsToCheck[higherPoints[l]] = true
				}
			}
		}
		pointsToCheck = newPointsToCheck
	}

	return basinPoints

}

func getSurroundingHigherPoints(grid [][]int, p point) []point {
	var higherPoints []point

	//Only points lower than 9 can have higher points surrounding it.
	if grid[p.i][p.j] < 9 {
		if p.i > 0 && grid[p.i-1][p.j] > grid[p.i][p.j] && grid[p.i-1][p.j] < 9 {
			pointToAdd := point{p.i - 1, p.j}
			higherPoints = append(higherPoints, pointToAdd)
		}
		if p.i < len(grid)-1 && grid[p.i+1][p.j] > grid[p.i][p.j] && grid[p.i+1][p.j] < 9 {
			pointToAdd := point{p.i + 1, p.j}
			higherPoints = append(higherPoints, pointToAdd)
		}
		if p.j > 0 && grid[p.i][p.j-1] > grid[p.i][p.j] && grid[p.i][p.j-1] < 9 {
			pointToAdd := point{p.i, p.j - 1}
			higherPoints = append(higherPoints, pointToAdd)
		}
		if p.j < len(grid[0])-1 && grid[p.i][p.j+1] > grid[p.i][p.j] && grid[p.i][p.j+1] < 9 {
			pointToAdd := point{p.i, p.j + 1}
			higherPoints = append(higherPoints, pointToAdd)
		}
	}
	return higherPoints
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
