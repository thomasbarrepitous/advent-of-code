package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type SurroundingStr struct {
	Symbol string
	X      int
	Y      int
}

var symbols = []string{"*"}

func populateSurroundingsSlice(inputMatrix [][]string, i int, j int) []SurroundingStr {
	surroundingsMatrix := make([]SurroundingStr, 0)
	if i != 0 {
		gear := SurroundingStr{Symbol: string(inputMatrix[i-1][j]), X: i - 1, Y: j}
		surroundingsMatrix = append(surroundingsMatrix, gear)
		if j != 0 {
			gear := SurroundingStr{Symbol: string(inputMatrix[i-1][j-1]), X: i - 1, Y: j - 1}
			surroundingsMatrix = append(surroundingsMatrix, gear)
		}
		if j != len(inputMatrix[i])-1 {
			gear := SurroundingStr{Symbol: string(inputMatrix[i-1][j+1]), X: i - 1, Y: j + 1}
			surroundingsMatrix = append(surroundingsMatrix, gear)
		}
	}
	if j != 0 {
		gear := SurroundingStr{Symbol: string(inputMatrix[i][j-1]), X: i, Y: j - 1}
		surroundingsMatrix = append(surroundingsMatrix, gear)
		if i != len(inputMatrix)-1 {
			gear := SurroundingStr{Symbol: string(inputMatrix[i+1][j-1]), X: i + 1, Y: j - 1}
			surroundingsMatrix = append(surroundingsMatrix, gear)
		}
	}
	if j != len(inputMatrix[i])-1 {
		gear := SurroundingStr{Symbol: string(inputMatrix[i][j+1]), X: i, Y: j + 1}
		surroundingsMatrix = append(surroundingsMatrix, gear)
		if i != len(inputMatrix)-1 {
			gear := SurroundingStr{Symbol: string(inputMatrix[i+1][j+1]), X: i + 1, Y: j + 1}
			surroundingsMatrix = append(surroundingsMatrix, gear)
		}
	}
	if i != len(inputMatrix)-1 {
		gear := SurroundingStr{Symbol: string(inputMatrix[i+1][j]), X: i + 1, Y: j}
		surroundingsMatrix = append(surroundingsMatrix, gear)
	}
	return surroundingsMatrix
}

func isValidCheck(inputMatrix [][]string, i int, j int, detectedInt *int, detectedGear *SurroundingStr) bool {
	// If next character is not ouf of bound
	if j+1 <= len(inputMatrix[i])-1 {
		parsedInt, err := strconv.Atoi(inputMatrix[i][j+1])
		// And is a valid integer
		if err == nil {
			*detectedInt = *detectedInt*10 + parsedInt
			// If valid we stop the recursion
			if isValidCheck(inputMatrix, i, j+1, detectedInt, detectedGear) {
				return true
			}
		}
	}
	// Check validity
	surroundingsSlice := populateSurroundingsSlice(inputMatrix, i, j)
	for _, surrouding := range surroundingsSlice {
		if slices.Contains(symbols, surrouding.Symbol) {
			*detectedGear = surrouding
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
	// Print the lines or perform other operations as needed
	inputMatrix := make([][]string, 0)
	for _, line := range lines {
		inputMatrix = append(inputMatrix, strings.Split(line, ""))
	}
	var flag int = 0
	var detectedInt int
	var detectedGear SurroundingStr
	detectedGearMap := map[SurroundingStr][]int{}
	for i := 0; i < len(inputMatrix); i++ {
		for j := 0; j < len(inputMatrix[i]); j++ {
			parsedInt, err := strconv.Atoi(inputMatrix[i][j])
			if err != nil {
				continue
			}
			detectedInt = parsedInt
			if isValidCheck(inputMatrix, i, j, &detectedInt, &detectedGear) {
				gearKey, ok := detectedGearMap[detectedGear]
				if ok {
					detectedGearMap[detectedGear] = append(gearKey, detectedInt)
				} else {
					detectedGearMap[detectedGear] = []int{detectedInt}
				}
				nbLength := len(strconv.Itoa(detectedInt))
				j += nbLength - 1
				detectedInt = 0
				detectedGear = SurroundingStr{}
			}
		}
	}
	for _, gear := range detectedGearMap {
		if len(gear) > 1 {
			flag += gear[0] * gear[1]
		}
	}
	// fmt.Println(detectedGearMap)
	fmt.Println(flag)
}

// Lesson learned: Use Row and Column instead of i and j
