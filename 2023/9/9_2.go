package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isHistoryOnlyZero(history []int) bool {
	for _, h := range history {
		if h != 0 {
			return false
		}
	}
	return true
}

func computeNextValue(history []int) int {
	processedHistory := make([]int, len(history)-1)
	for i := 0; i < len(history)-1; i++ {
		firstTerm := history[i]
		secondTerm := history[i+1]
		processedHistory[i] = secondTerm - firstTerm
	}
	// fmt.Println(processedHistory)
	if isHistoryOnlyZero(processedHistory) {
		return 0
	}
	return processedHistory[0] - computeNextValue(processedHistory)
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
	for _, line := range lines {
		var history []int
		// fmt.Println(line)
		linesSplit := strings.Split(line, " ")
		for _, value := range linesSplit {
			parsedValue, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Error parsing value:", err)
				return
			}
			history = append(history, parsedValue)
		}
		// fmt.Println(history[0] - computeNextValue(history))
		flag += history[0] - computeNextValue(history)
	}

	fmt.Println("flag:", flag)
}
