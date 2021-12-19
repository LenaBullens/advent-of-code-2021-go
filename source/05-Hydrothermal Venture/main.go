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

type line struct {
	p1 point
	p2 point
}

func main() {
	lines := readInput()
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(lines []line) int {
	cloudMap := map[point]int{}

	for i := 0; i < len(lines); i++ {
		pointsOfLine := calculatePointsInLine(lines[i], false)
		for j := 0; j < len(pointsOfLine); j++ {
			//Key equality is checked using '=='. Two structs are equal if they are the same type and all their
			//fields are equal.
			cloudMap[pointsOfLine[j]] = cloudMap[pointsOfLine[j]] + 1
		}
	}

	result := 0

	for _, count := range cloudMap {
		if count >= 2 {
			result = result + 1
		}
	}
	return result
}

func part2(lines []line) int {
	cloudMap := map[point]int{}

	for i := 0; i < len(lines); i++ {
		pointsOfLine := calculatePointsInLine(lines[i], true)
		for j := 0; j < len(pointsOfLine); j++ {
			//Key equality is checked using '=='. Two structs are equal if they are the same type and all their
			//fields are equal.
			cloudMap[pointsOfLine[j]] = cloudMap[pointsOfLine[j]] + 1
		}
	}

	result := 0

	for _, count := range cloudMap {
		if count >= 2 {
			result = result + 1
		}
	}
	return result
}

func readInput() []line {
	f, err := os.Open("input-05.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []line

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		strPoint1 := strings.Split(split[0], ",")
		x1, err := strconv.Atoi(strPoint1[0])
		if err != nil {
			log.Fatal(err)
		}
		y1, err := strconv.Atoi(strPoint1[1])
		if err != nil {
			log.Fatal(err)
		}
		point1 := point{x1, y1}
		strPoint2 := strings.Split(split[2], ",")
		x2, err := strconv.Atoi(strPoint2[0])
		if err != nil {
			log.Fatal(err)
		}
		y2, err := strconv.Atoi(strPoint2[1])
		if err != nil {
			log.Fatal(err)
		}
		point2 := point{x2, y2}
		line := line{point1, point2}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func calculatePointsInLine(line line, withDiagonals bool) []point {
	var points []point
	if isHorizontalLine(line) {
		if line.p1.x < line.p2.x {
			for i := line.p1.x; i <= line.p2.x; i++ {
				points = append(points, point{i, line.p1.y})
			}
		} else {
			for i := line.p2.x; i <= line.p1.x; i++ {
				points = append(points, point{i, line.p2.y})
			}
		}
	} else if isVerticalLine(line) {
		if line.p1.y < line.p2.y {
			for i := line.p1.y; i <= line.p2.y; i++ {
				points = append(points, point{line.p1.x, i})
			}
		} else {
			for i := line.p2.y; i <= line.p1.y; i++ {
				points = append(points, point{line.p2.x, i})
			}
		}
	} else if withDiagonals {
		//A diagonal goes either 'left' or 'right' and 'up' or 'down'; so there are four different cases.
		//Horizonal and vertical difference in coordinates will be the same because all diagonals are at 45 degrees.
		if line.p1.x < line.p2.x {
			if line.p1.y < line.p2.y {
				for i := 0; i <= line.p2.x-line.p1.x; i++ {
					points = append(points, point{line.p1.x + i, line.p1.y + i})
				}
			} else {
				for i := 0; i <= line.p2.x-line.p1.x; i++ {
					points = append(points, point{line.p1.x + i, line.p1.y - i})
				}
			}
		} else {
			if line.p1.y < line.p2.y {
				for i := 0; i <= line.p1.x-line.p2.x; i++ {
					points = append(points, point{line.p1.x - i, line.p1.y + i})
				}
			} else {
				for i := 0; i <= line.p1.x-line.p2.x; i++ {
					points = append(points, point{line.p1.x - i, line.p1.y - i})
				}
			}
		}
	}
	return points
}

func isHorizontalLine(line line) bool {
	return line.p1.y == line.p2.y
}

func isVerticalLine(line line) bool {
	return line.p1.x == line.p2.x
}
