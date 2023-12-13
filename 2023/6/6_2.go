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
	var flag int = 0
	str_time_inputs := strings.Split(lines[0], ":")
	str_distance_line := strings.Split(lines[1], ":")

	time_input, err := strconv.Atoi(strings.ReplaceAll(str_time_inputs[1], " ", ""))
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	distance_input, err := strconv.Atoi(strings.ReplaceAll(str_distance_line[1], " ", ""))
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(lines)
	fmt.Println(time_input)
	fmt.Println(distance_input)
	inputs := []Input{{time_input, distance_input}}
	for _, input := range inputs {
		for speed := 0; speed < input.Time+1; speed++ {
			distance := speed * (input.Time - speed)
			if distance > input.Distance {
				flag++
			}
		}
		fmt.Println("Input: ", input)
	}
	fmt.Println(flag)
}
