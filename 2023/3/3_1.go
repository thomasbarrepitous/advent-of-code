package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var symbols = []string{"-", "#", "*", "%", "&", "$", "@", "=", "/", "+"}

func populateSurroundingsMatrix(inputMatrix [][]string, i int, j int) []string {
	surroundingsMatrix := make([]string, 0)
	if i != 0 {
		surroundingsMatrix = append(surroundingsMatrix, inputMatrix[i-1][j])
		if j != 0 {
			surroundingsMatrix = append(surroundingsMatrix, inputMatrix[i-1][j-1])
		}
		if j != len(inputMatrix[i])-1 {
			surroundingsMatrix = append(surroundingsMatrix, inputMatrix[i-1][j+1])
		}
	}
	if j != 0 {
		surroundingsMatrix = append(surroundingsMatrix, inputMatrix[i][j-1])
		if i != len(inputMatrix)-1 {
			surroundingsMatrix = append(surroundingsMatrix, inputMatrix[i+1][j-1])
		}
	}
	if j != len(inputMatrix[i])-1 {
		surroundingsMatrix = append(surroundingsMatrix, inputMatrix[i][j+1])
		if i != len(inputMatrix)-1 {
			surroundingsMatrix = append(surroundingsMatrix, inputMatrix[i+1][j+1])
		}
	}
	if i != len(inputMatrix)-1 {
		surroundingsMatrix = append(surroundingsMatrix, inputMatrix[i+1][j])
	}
	return surroundingsMatrix
}

func isValidCheck(inputMatrix [][]string, i int, j int, val *int) bool {
	// If next character is not ouf of bound
	if j+1 <= len(inputMatrix[i])-1 {
		// And is a valid integer
		parsedInt, err := strconv.Atoi(inputMatrix[i][j+1])
		if err == nil {
			*val = *val*10 + parsedInt
			// If valid we stop the recursion
			if isValidCheck(inputMatrix, i, j+1, val) {
				return true
			}
		}
	}
	// Check validity
	surroundingsMatrix := populateSurroundingsMatrix(inputMatrix, i, j)
	for _, surrouding := range surroundingsMatrix {
		if slices.Contains(symbols, surrouding) {
			return true
		}
	}
	return false
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
	// Print the lines or perform other operations as needed
	inputMatrix := make([][]string, 0)
	for _, line := range lines {
		inputMatrix = append(inputMatrix, strings.Split(line, ""))
	}
	var detectedInt int
	m := make(map[string]string)
	for i := 0; i < len(inputMatrix); i++ {
		for j := 0; j < len(inputMatrix[i]); j++ {
			// fmt.Println(inputMatrix[i][j])
			parsedInt, err := strconv.Atoi(inputMatrix[i][j])
			if err != nil {
				if inputMatrix[i][j] != "." {
					m[inputMatrix[i][j]] = inputMatrix[i][j]
				}
				continue
			}
			detectedInt = parsedInt
			if isValidCheck(inputMatrix, i, j, &detectedInt) {
				flag += detectedInt
				// fmt.Println(detectedInt)
				nbLength := len(strconv.Itoa(detectedInt))
				j += nbLength - 1
				detectedInt = 0
			}
		}
	}
	fmt.Println(m)
	fmt.Println(flag)
}
