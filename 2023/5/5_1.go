package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type InputRange struct {
	destination int
	source      int
	length      int
}

func fromStringArrayToInputArray(inputArray []string) (input []InputRange) {
	for _, inputStr := range inputArray {
		inputSplit := strings.Split(inputStr, " ")
		destination, err := strconv.Atoi(inputSplit[0])
		if err != nil {
			continue
		}
		source, err := strconv.Atoi(inputSplit[1])
		if err != nil {
			continue
		}
		rangeLen, err := strconv.Atoi(inputSplit[2])
		if err != nil {
			continue
		}
		input = append(input, InputRange{destination, source, rangeLen})
	}
	return
}

func main() {
	// Specify the path to your text file
	filePath := "input.txt"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a slice to store the lines
	var lines []string

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	flags := []int{}
	var inputArray []string
	for _, line := range lines {
		inputArray = append(inputArray, line)
	}

	seedsSplit := strings.Split(inputArray[0], ":")
	seedsInput := strings.Split(seedsSplit[1], " ")

	for _, seedStr := range seedsInput {
		seed, err := strconv.Atoi(string(seedStr))
		if err != nil {
			continue
		}
		flags = append(flags, seed)
	}

	// Not very pretty isn't it?
	seedToSoilLines := inputArray[2:30]
	soilToFertilizerLines := inputArray[31:52]
	fertilizerToWaterLines := inputArray[53:102]
	waterToLightLines := inputArray[103:146]
	lightToTemperatureLines := inputArray[147:172]
	temperatureToHumidityLines := inputArray[173:199]
	humidityToLocationLines := inputArray[200:225]
	slicedInputs := [][]string{seedToSoilLines, soilToFertilizerLines, fertilizerToWaterLines, waterToLightLines, lightToTemperatureLines, temperatureToHumidityLines, humidityToLocationLines}
	var parsedInput [][]InputRange
	for _, slicedInput := range slicedInputs {
		parsedInput = append(parsedInput, fromStringArrayToInputArray(slicedInput))
	}

	for idx := 0; idx < len(flags); idx++ {
	NEXT_DESTINATION:
		for _, inputs := range parsedInput {
			for _, input := range inputs {
				if flags[idx] >= input.source && flags[idx] <= input.source+input.length {
					flags[idx] = input.destination - input.source + flags[idx]
					continue NEXT_DESTINATION
				}
			}
		}
	}
	slices.Sort(flags)
	fmt.Println(flags[0])
}
