package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const MAXINT = int(^uint(0) >> 1)

type point struct {
	row    int
	column int
}

func main() {
	part2()
}

func part1() {
	grid := readInput("input-15.txt")
	height := len(grid)
	width := len(grid[0])

	unvisitedNodes := make(map[point]bool)
	tentativeDistances := make(map[point]int)

	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if r != 0 || c != 0 {
				p := createPoint(r, c)
				unvisitedNodes[p] = true
				tentativeDistances[p] = MAXINT
			}
		}
	}

	start := createPoint(0, 0)
	end := createPoint(height-1, width-1)
	tentativeDistances[start] = 0
	done := false

	for !done {
		unvisitedNodes, tentativeDistances = step(start, unvisitedNodes, tentativeDistances, grid, height, width)

		//Check if visited end
		_, exists := unvisitedNodes[end]
		if !exists {
			done = true
		} else {
			start = findLowestUnvisitedNode(unvisitedNodes, tentativeDistances)
		}
	}
	fmt.Println(tentativeDistances[end])
}

func part2() {
	grid := expandGrid(readInput("input-15.txt"))
	height := len(grid)
	width := len(grid[0])

	unvisitedNodes := make(map[point]bool)
	tentativeDistances := make(map[point]int)

	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if r != 0 || c != 0 {
				p := createPoint(r, c)
				unvisitedNodes[p] = true
				//Only keep track of neighbouring notes distances to speed up finding lowest node.
				//tentativeDistances[p] = MAXINT
			}
		}
	}

	start := createPoint(0, 0)
	end := createPoint(height-1, width-1)
	tentativeDistances[start] = 0
	done := false
	counter := 0

	for !done {
		fmt.Print("Checking point: (")
		fmt.Print(start.row)
		fmt.Print(",")
		fmt.Print(start.column)
		fmt.Print(") - ")
		fmt.Print(counter)
		fmt.Println(" points checked.")
		unvisitedNodes, tentativeDistances = step2(start, unvisitedNodes, tentativeDistances, grid, height, width)

		//Check if visited end
		_, exists := unvisitedNodes[end]
		if !exists {
			done = true
		} else {
			start = findLowestUnvisitedNode2(unvisitedNodes, tentativeDistances)
			counter++
		}
	}
	fmt.Println(tentativeDistances[end])
}

func step(node point, unvisitedNodes map[point]bool, tentativeDistances map[point]int, grid [][]int, height int, width int) (map[point]bool, map[point]int) {
	neighbours := findUnvisitedNeighbours(node, unvisitedNodes, height, width)

	startDistance := tentativeDistances[node]
	for _, n := range neighbours {
		newDistance := startDistance + grid[n.row][n.column]
		if newDistance < tentativeDistances[n] {
			tentativeDistances[n] = newDistance
		}
	}
	delete(unvisitedNodes, node)
	return unvisitedNodes, tentativeDistances
}

func step2(node point, unvisitedNodes map[point]bool, tentativeDistances map[point]int, grid [][]int, height int, width int) (map[point]bool, map[point]int) {
	neighbours := findUnvisitedNeighbours(node, unvisitedNodes, height, width)

	startDistance := tentativeDistances[node]
	for _, n := range neighbours {
		newDistance := startDistance + grid[n.row][n.column]
		//We're only keeping track of distances for nodes that have been neighbours of a
		//node being visited, others will always be at max distance. This will hopefully
		//speed up finding our next lowest node.
		_, exists := tentativeDistances[n]
		if exists {
			if newDistance < tentativeDistances[n] {
				tentativeDistances[n] = newDistance
			}
		} else {
			tentativeDistances[n] = newDistance
		}

	}
	delete(unvisitedNodes, node)
	return unvisitedNodes, tentativeDistances
}

func findLowestUnvisitedNode(unvisitedNodes map[point]bool, tentativeDistances map[point]int) point {
	minimum := MAXINT
	var result point
	for key, _ := range unvisitedNodes {
		tDis := tentativeDistances[key]
		if tDis < minimum {
			minimum = tDis
			result = key
		}
	}
	return result
}

func findLowestUnvisitedNode2(unvisitedNodes map[point]bool, tentativeDistances map[point]int) point {
	minimum := MAXINT
	var result point
	for key, value := range tentativeDistances {
		_, exists := unvisitedNodes[key]
		if exists {
			if value < minimum {
				minimum = value
				result = key
			}
		}
	}
	return result
}

func findUnvisitedNeighbours(p point, unvisitedNodes map[point]bool, height int, width int) []point {
	var neighbours []point
	//Left neighbour
	if p.column > 0 {
		neighbour := createPoint(p.row, p.column-1)
		_, exists := unvisitedNodes[neighbour]
		if exists {
			neighbours = append(neighbours, neighbour)
		}
	}
	//Top neighbour
	if p.row > 0 {
		neighbour := createPoint(p.row-1, p.column)
		_, exists := unvisitedNodes[neighbour]
		if exists {
			neighbours = append(neighbours, neighbour)
		}
	}
	//Right neighbour
	if p.column < width {
		neighbour := createPoint(p.row, p.column+1)
		_, exists := unvisitedNodes[neighbour]
		if exists {
			neighbours = append(neighbours, neighbour)
		}
	}
	//Bottom neighbour
	if p.row < height {
		neighbour := createPoint(p.row+1, p.column)
		_, exists := unvisitedNodes[neighbour]
		if exists {
			neighbours = append(neighbours, neighbour)
		}
	}
	return neighbours
}

func createPoint(row int, column int) point {
	p := point{}
	p.row = row
	p.column = column
	return p
}

func readInput(path string) [][]int {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var grid [][]int

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var gridRow []int
		splitLine := strings.Split(scanner.Text(), "")
		for i := 0; i < len(splitLine); i++ {
			risk, err := strconv.Atoi(splitLine[i])
			if err != nil {
				log.Fatal(err)
			}
			gridRow = append(gridRow, risk)
		}
		grid = append(grid, gridRow)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid
}

func expandGrid(initialGrid [][]int) [][]int {
	height := len(initialGrid)
	width := len(initialGrid[0])
	result := make([][]int, height*5)
	for i := 0; i < height*5; i++ {
		result[i] = make([]int, width*5)
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for r := 0; r < height; r++ {
				for c := 0; c < width; c++ {
					newRisk := initialGrid[r][c] + i + j
					if newRisk > 9 {
						newRisk = newRisk - 9
					}
					result[r+height*i][c+width*j] = newRisk
				}
			}
		}
	}
	return result
}
