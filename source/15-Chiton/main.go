package main

import (
	"bufio"
	"container/heap"
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

type Item struct {
	value    point
	priority int
	index    int
	visited  bool
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value point, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	part1()
	part2()
}

func part1() {
	solve(false)
}

func part2() {
	solve(true)
}

func solve(expand bool) {
	grid := readInput("input-15.txt")
	if expand {
		grid = expandGrid(grid)
	}
	height := len(grid)
	width := len(grid[0])

	tentativeDistances := make(map[*Item]int, height*width)
	priorityQueue := make(PriorityQueue, height*width)
	chitons := make(map[point]*Item, height*width)

	i := 0
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if r == 0 && c == 0 {
				p := createPoint(r, c)
				item := Item{value: p, priority: 0, index: i, visited: false}
				tentativeDistances[&item] = 0
				priorityQueue[i] = &item
				chitons[p] = &item
			} else {
				p := createPoint(r, c)
				item := Item{value: p, priority: MAXINT, index: i, visited: false}
				tentativeDistances[&item] = MAXINT
				priorityQueue[i] = &item
				chitons[p] = &item
			}
			i = i + 1
		}
	}

	heap.Init(&priorityQueue)

	for priorityQueue.Len() > 0 {
		nextNode := heap.Pop(&priorityQueue).(*Item)
		unvisitedNeighbours := findUnvisitedNeighbours(nextNode, chitons, height, width)
		for i := 0; i < len(unvisitedNeighbours); i++ {
			row := unvisitedNeighbours[i].value.row
			column := unvisitedNeighbours[i].value.column
			newDistance := tentativeDistances[nextNode] + grid[row][column]
			if newDistance < tentativeDistances[unvisitedNeighbours[i]] {
				tentativeDistances[unvisitedNeighbours[i]] = newDistance
				priorityQueue.update(unvisitedNeighbours[i], unvisitedNeighbours[i].value, newDistance)
			}
		}
		nextNode.visited = true
	}
	end := point{row: height - 1, column: width - 1}
	endItem := chitons[end]
	result := tentativeDistances[endItem]
	fmt.Printf("Result: %d\n", result)
}

func findUnvisitedNeighbours(item *Item, chitons map[point]*Item, height int, width int) []*Item {
	var neighbours []*Item
	//Left neighbour
	if item.value.column > 0 {
		neighbour := createPoint(item.value.row, item.value.column-1)
		_, exists := chitons[neighbour]
		if exists {
			neighbours = append(neighbours, chitons[neighbour])
		}
	}
	//Top neighbour
	if item.value.row > 0 {
		neighbour := createPoint(item.value.row-1, item.value.column)
		_, exists := chitons[neighbour]
		if exists && !chitons[neighbour].visited {
			neighbours = append(neighbours, chitons[neighbour])
		}
	}
	//Right neighbour
	if item.value.column < width {
		neighbour := createPoint(item.value.row, item.value.column+1)
		_, exists := chitons[neighbour]
		if exists && !chitons[neighbour].visited {
			neighbours = append(neighbours, chitons[neighbour])
		}
	}
	//Bottom neighbour
	if item.value.row < height {
		neighbour := createPoint(item.value.row+1, item.value.column)
		_, exists := chitons[neighbour]
		if exists && !chitons[neighbour].visited {
			neighbours = append(neighbours, chitons[neighbour])
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
