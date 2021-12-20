package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type target struct {
	x1 int
	x2 int
	y1 int
	y2 int
}

func main() {
	target := readInput("input-17.txt")
	solve(target)
}

func solve(target target) {
	var minXVelocity int
	var maxXVelocity int
	var minYVelocity int
	var maxYVelocity int

	//Velocity in the X direction can't be bigger than the right most bound of the target; otherwise
	//we will already overshoot the target on step 1.
	maxXVelocity = target.x2

	//Total distance reached is triangular number of the starting horizontal velocity. Can solve as
	//quadratic equation to find lower bound for velocity.
	minXVelocity = determineMinXVelocity(target.x1)

	//Vertical speed can't be smaller than lowest bound (assuming lowest bound negative). Because in that
	//case the probe would already be lower after 1 step.
	minYVelocity = target.y1

	//Shoot to just hit bottom. Max speed positive. Pas y = 0 again at speed v0-1, step later are at y = v0-1.
	//lower bound = v0 - 1 => v0 = -lower bound -1
	maxYVelocity = -1*target.y1 - 1

	highestY := 0
	counter := 0

	for i := minXVelocity; i <= maxXVelocity; i++ {
		for j := minYVelocity; j <= maxYVelocity; j++ {
			done := false
			maxY := 0
			hitTarget := false
			vx := i
			vy := j
			x := 0
			y := 0
			for !done {
				x = x + vx
				y = y + vy
				if vx > 0 {
					vx = vx - 1
				}
				vy = vy - 1
				if y > maxY {
					maxY = y
				}
				if x >= target.x1 && x <= target.x2 && y >= target.y1 && y <= target.y2 {
					hitTarget = true
				}
				if x > target.x2 || y < target.y1 {
					done = true
				}
			}
			if hitTarget {
				counter++
				if maxY > highestY {
					highestY = maxY
				}
			}
		}
	}

	fmt.Println(highestY)
	fmt.Println(counter)
}

func determineMinXVelocity(lowerbound int) int {
	//Add one before converting back to int to round up
	return int((-1+math.Sqrt(float64(1+8*lowerbound)))/2 + 1)
}

func readInput(path string) target {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var result target

	for scanner.Scan() {
		result = target{}
		splitString := strings.Split(scanner.Text()[15:], ",")
		splitXString := strings.Split(splitString[0], "..")
		splitYString := strings.Split(splitString[1][3:], "..")
		x1, err := strconv.Atoi(splitXString[0])
		if err != nil {
			log.Fatal(err)
		}
		x2, err := strconv.Atoi(splitXString[1])
		if err != nil {
			log.Fatal(err)
		}
		y1, err := strconv.Atoi(splitYString[0])
		if err != nil {
			log.Fatal(err)
		}
		y2, err := strconv.Atoi(splitYString[1])
		if err != nil {
			log.Fatal(err)
		}
		result.x1 = x1
		result.x2 = x2
		result.y1 = y1
		result.y2 = y2
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
