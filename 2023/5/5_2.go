package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
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

func getLocationFromSeed(seed int, parsedInput [][]InputRange) (location int) {
	source := seed
NEXT_DESTINATION:
	for _, inputs := range parsedInput {
		for _, input := range inputs {
			if source >= input.source && source <= input.source+input.length {
				source = input.destination - input.source + source
				continue NEXT_DESTINATION
			}
		}
	}
	return source
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
	flag := 0
	seeds := []int{}
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
		seeds = append(seeds, seed)
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

	flagChan := make(chan int, 1000)
	wg := &sync.WaitGroup{}
	for idx := 0; idx < len(seeds); idx += 2 {
		baseSeed := seeds[idx]
		seedRange := seeds[idx+1]
		go func() {
			for location := range flagChan {
				if location < flag || flag == 0 {
					flag = location
				}
				wg.Done()
			}
		}()

		for seed := baseSeed; seed < baseSeed+seedRange; seed += 1 {
			wg.Add(1)
			go func(seed int) {
				flagChan <- getLocationFromSeed(seed, parsedInput)
			}(seed)
			fmt.Printf("\r Seeds: %d/%d, Seed: %d/%d, Flag: %d", idx/2+1, len(seeds)/2+1, (seed - baseSeed), seedRange, flag)
		}
	}
	wg.Wait()
	close(flagChan)
	fmt.Printf("FLAG : %d", flag)
}
