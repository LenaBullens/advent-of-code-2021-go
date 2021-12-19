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

var hexToBinMap = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}

func main() {
	hex := readInput("input-16.txt")
	bin := hexToBin(hex)
	totalVersion, totalValue, _ := parsePacket(bin)
	fmt.Println("==========")
	fmt.Printf("Total version: %d\n", totalVersion)
	fmt.Printf("Total value: %d\n", totalValue)
}

func parsePacket(bin string) (int, int, int) {
	//Minimal length of a packet is 11; a literal packet with 4 bit of literal data.
	if len(bin) < 11 {
		return 0, 0, 0
	}

	//Amount of bits parsed in this package.
	bitsParsed := 0
	packetValue := 0

	//Version
	packetVersion := binToInt(bin[bitsParsed : bitsParsed+3])
	bitsParsed = bitsParsed + 3

	//Type
	packetType := binToInt(bin[bitsParsed : bitsParsed+3])
	bitsParsed = bitsParsed + 3

	if packetType == 4 {
		doneLiteral := false
		var valueInBin string
		for !doneLiteral {
			fiveBits := bin[bitsParsed : bitsParsed+5]
			oneBit := fiveBits[:1]
			valueInBin = valueInBin + fiveBits[1:]
			bitsParsed = bitsParsed + 5
			if oneBit == "0" {
				doneLiteral = true
			}
		}
		packetValue = binToInt(valueInBin)
	} else {
		//Check type of length
		lengthType := bin[bitsParsed : bitsParsed+1]
		bitsParsed = bitsParsed + 1
		if lengthType == "0" {
			length := binToInt(bin[bitsParsed : bitsParsed+15])
			bitsParsed = bitsParsed + 15
			subBitsParsed := 0
			var values []int
			for subBitsParsed < length {
				subVersion, subValue, subBits := parsePacket(bin[bitsParsed+subBitsParsed:])
				packetVersion = packetVersion + subVersion
				values = append(values, subValue)
				subBitsParsed = subBitsParsed + subBits
			}
			bitsParsed = bitsParsed + subBitsParsed
			packetValue = calculateValue(values, packetType)
		} else {
			nbOfPackets := binToInt(bin[bitsParsed : bitsParsed+11])
			bitsParsed = bitsParsed + 11
			subBitsParsed := 0
			var values []int
			for i := 0; i < nbOfPackets; i++ {
				subVersion, subValue, subBits := parsePacket(bin[bitsParsed+subBitsParsed:])
				packetVersion = packetVersion + subVersion
				values = append(values, subValue)
				subBitsParsed = subBitsParsed + subBits
			}
			bitsParsed = bitsParsed + subBitsParsed
			packetValue = calculateValue(values, packetType)
		}
	}
	return packetVersion, packetValue, bitsParsed
}

func calculateValue(values []int, packetType int) int {
	result := 0
	if packetType == 0 {
		for _, v := range values {
			result = result + v
		}
	} else if packetType == 1 {
		result = 1
		for _, v := range values {
			result = result * v
		}
	} else if packetType == 2 {
		min := math.MaxInt32
		for _, v := range values {
			if v < min {
				min = v
			}
		}
		result = min
	} else if packetType == 3 {
		max := math.MinInt32
		for _, v := range values {
			if v > max {
				max = v
			}
		}
		result = max
	} else if packetType == 5 {
		if values[0] > values[1] {
			result = 1
		}
	} else if packetType == 6 {
		if values[0] < values[1] {
			result = 1
		}
	} else if packetType == 7 {
		if values[0] == values[1] {
			result = 1
		}
	}
	return result
}

func hexToBin(hex string) string {
	var result string
	splitString := strings.Split(hex, "")
	for i := 0; i < len(splitString); i++ {
		result = result + hexToBinMap[splitString[i]]
	}
	return result
}

func binToInt(bin string) int {
	result, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(result)
}

func readInput(path string) string {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var result string

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		result = result + scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
