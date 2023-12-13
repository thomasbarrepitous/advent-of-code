package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Time     int
	Distance int
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
	var flag int = 1
	str_time_inputs := strings.Split(lines[0], ":")
	str_distance_line := strings.Split(lines[1], ":")
	raw_time_inputs := strings.Split(str_time_inputs[1], " ")
	raw_distance_inputs := strings.Split(str_distance_line[1], " ")
	time_inputs := []int{}
	distance_inputs := []int{}
	inputs := []Input{}
	for _, timeStr := range raw_time_inputs {
		time, err := strconv.Atoi(timeStr)
		if err != nil {
			continue
		}
		time_inputs = append(time_inputs, time)
	}
	for _, distanceStr := range raw_distance_inputs {
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			continue
		}
		distance_inputs = append(distance_inputs, distance)
	}
	for idx, time := range time_inputs {
		inputs = append(inputs, Input{time, distance_inputs[idx]})
	}

	fmt.Println(lines)
	for _, input := range inputs {
		records := 0
		for speed := 0; speed < input.Time+1; speed++ {
			distance := speed * (input.Time - speed)
			if distance > input.Distance {
				records++
			}
		}
		fmt.Println("Input: ", input)
		fmt.Println("Records: ", records)
		flag = flag * records
	}
	fmt.Println(flag)
}
