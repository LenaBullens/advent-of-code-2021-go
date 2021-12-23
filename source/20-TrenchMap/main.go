package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type pixel struct {
	row    int
	column int
}

type image struct {
	pixels map[pixel]bool
	min    int
	max    int
}

var algorithm []bool

func main() {
	part1()
	part2()
}

func part1() {
	solve(2)
}

func part2() {
	solve(50)
}

func solve(times int) {
	algorithmString, imageStringSlice := readInput("input-20.txt")
	algorithm = parseAlgorithm(algorithmString)
	image := parseImage(imageStringSlice)
	background := false
	for i := 0; i < times; i++ {
		image = enhanceImage(image, background)
		//Adjust the background
		if background {
			//'111111111' is 511
			background = algorithm[511]
		} else {
			background = algorithm[0]
		}
	}

	count := 0
	for _, lit := range image.pixels {
		if lit {
			count = count + 1
		}
	}
	fmt.Printf("Number of lit pixels: %d\n", count)
}

func parseAlgorithm(algorithmString string) []bool {
	var result []bool
	for i := 0; i < len(algorithmString); i++ {
		boolValue := algorithmString[i:i+1] == "#"
		result = append(result, boolValue)
	}
	return result
}

func parseImage(imageStringSlice []string) image {
	pixels := make(map[pixel]bool)
	for r := 0; r < len(imageStringSlice); r++ {
		for c := 0; c < len(imageStringSlice[r]); c++ {
			pixel := pixel{}
			pixel.row = r
			pixel.column = c
			pixels[pixel] = imageStringSlice[r][c:c+1] == "#"
		}
	}
	min := 0
	max := len(imageStringSlice) - 1
	image := image{}
	image.pixels = pixels
	image.min = min
	image.max = max
	return image
}

func enhancePixel(p pixel, img image, background bool) bool {
	var binString string
	neighbours := calculateNeighbours(p)
	for _, n := range neighbours {
		lit, known := img.pixels[n]
		//If pixel is not known it's part of the background.
		if !known {
			lit = background
		}
		if lit {
			binString = binString + "1"
		} else {
			binString = binString + "0"
		}
	}
	return algorithm[binToInt(binString)]
}

func enhanceImage(img image, background bool) image {
	img = expandImage(img, background)
	newImg := image{}
	newImg.pixels = make(map[pixel]bool)
	newImg.min = img.min
	newImg.max = img.max
	for p := range img.pixels {
		newImg.pixels[p] = enhancePixel(p, img, background)
	}
	return newImg
}

func expandImage(img image, background bool) image {
	img.min = img.min - 1
	img.max = img.max + 1
	for r := img.min; r <= img.max; r++ {
		for c := img.min; c <= img.max; c++ {
			p := pixel{}
			p.row = r
			p.column = c
			_, known := img.pixels[p]
			if !known {
				img.pixels[p] = background
			}
		}
	}
	return img
}

func calculateNeighbours(p pixel) []pixel {
	var neighbours []pixel
	for r := p.row - 1; r <= p.row+1; r++ {
		for c := p.column - 1; c <= p.column+1; c++ {
			neighbour := pixel{}
			neighbour.row = r
			neighbour.column = c
			neighbours = append(neighbours, neighbour)
		}
	}
	return neighbours
}

func binToInt(bin string) int {
	result, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(result)
}

func readInput(path string) (string, []string) {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var algorithm string
	var image []string
	mode := 0

	for scanner.Scan() {
		line := scanner.Text()
		if mode == 0 {
			if line != "" {
				algorithm = algorithm + line
			} else {
				mode = 1
			}
		} else {
			image = append(image, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return algorithm, image
}
