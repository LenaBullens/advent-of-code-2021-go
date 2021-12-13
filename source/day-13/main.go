package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type fold struct {
	direction string
	value     int
}

func main() {
	part1()
	part2()
}

func part1() {
	points, folds := readInput("input-13.txt")

	//Do one fold
	//All points further from origin than the fold line get folded. New position (for the axis orthogonal to the fold line)
	//is a_fold - (a_point - a_fold) = 2*a_fold - a_point
	fold := folds[0]
	if fold.direction == "horizontal" {
		for p, _ := range points {
			if p.y > fold.value {
				newY := 2*fold.value - p.y
				newPoint := point{p.x, newY}
				delete(points, p)
				points[newPoint] = true
			}
		}
	} else {
		for p, _ := range points {
			if p.x > fold.value {
				newX := 2*fold.value - p.x
				newPoint := point{newX, p.y}
				delete(points, p)
				points[newPoint] = true
			}
		}
	}

	totalPoints := len(points)
	fmt.Println(totalPoints)
}

func part2() {
	points, folds := readInput("input-13.txt")

	for i := 0; i < len(folds); i++ {
		fold := folds[i]
		if fold.direction == "horizontal" {
			for p, _ := range points {
				if p.y > fold.value {
					newY := 2*fold.value - p.y
					newPoint := point{p.x, newY}
					delete(points, p)
					points[newPoint] = true
				}
			}
		} else {
			for p, _ := range points {
				if p.x > fold.value {
					newX := 2*fold.value - p.x
					newPoint := point{newX, p.y}
					delete(points, p)
					points[newPoint] = true
				}
			}
		}
	}

	maxX, maxY := findMaxXY(points)
	for i := 0; i <= maxY; i++ {
		for j := 0; j <= maxX; j++ {
			p := point{j, i}
			if points[p] {
				fmt.Print("# ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println("")
	}
}

func findMaxXY(points map[point]bool) (int, int) {
	var maxX int = -1
	var maxY int = -1
	for p, _ := range points {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	return maxX, maxY
}

func readInput(path string) (map[point]bool, []fold) {
	mode := 0 //0 is scanning points, 1 is scanning folds
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	points := make(map[point]bool)
	var folds []fold

	for scanner.Scan() {
		line := scanner.Text()
		if mode == 0 {
			if line != "" {
				splitString := strings.Split(line, ",")
				x, err := strconv.Atoi(splitString[0])
				if err != nil {
					log.Fatal(err)
				}
				y, err := strconv.Atoi(splitString[1])
				if err != nil {
					log.Fatal(err)
				}
				p := point{x, y}
				points[p] = true
			} else {
				mode = 1
			}
		} else if mode == 1 {
			dataString := line[11:]
			splitString := strings.Split(dataString, "=")
			value, err := strconv.Atoi(splitString[1])
			if err != nil {
				log.Fatal(err)
			}
			var direction string
			if splitString[0] == "y" {
				direction = "horizontal"
			} else {
				direction = "vertical"
			}
			f := fold{direction, value}
			folds = append(folds, f)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return points, folds
}
