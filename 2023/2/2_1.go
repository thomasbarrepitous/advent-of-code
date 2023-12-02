package main


import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "strconv"
    "log"
)

const RED_LIMIT int = 12
const GREEN_LIMIT int = 13
const BLUE_LIMIT int = 14

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
	for _, line := range lines {
        fmt.Println(line)
        // Delimiter game nb / input
        lineSplit := strings.Split(line, ":")
        // Get game nb
        gameSplit:= strings.Split(lineSplit[0], "Game")
        gameNb, err := strconv.Atoi(strings.TrimSpace(gameSplit[1]))
        if err != nil {
            log.Fatal(err)
        }
        if isGameValidChecker(lineSplit[1]){
            flag += gameNb
        }
	}
    fmt.Println(flag)
}

func colourChecker(res string) string{
    if strings.Contains(res, "red") {
        return "red"
    }
    if strings.Contains(res , "green") {
        return "green"
    }
    return "blue"
}

func isCubeValidChecker(cube string) bool{
    colour := colourChecker(cube)
    cubeSplit:= strings.Split(cube, colour)
    cubeValue, err := strconv.Atoi(strings.TrimSpace(cubeSplit[0]))
    if err != nil {
        log.Fatal(err)
    }
    if colour == "red" && cubeValue > RED_LIMIT {
        return false
    }
    if colour == "green" && cubeValue > GREEN_LIMIT {
        return false
    }
    if colour == "blue" && cubeValue > BLUE_LIMIT {
        return false
    }
    return true
}

func isRoundValidChecker(gameResults string) bool{
    // Delimiter for each cube
    roundResults := strings.Split(gameResults, ",")
    for _, cube := range roundResults{
        if !isCubeValidChecker(cube){
            return false
        }
    }
    return true
}

func isGameValidChecker(gameInput string) bool{
    // Delimiter for each round
    gameInputSplit:= strings.Split(gameInput, ";")
    for _, gameResults:= range gameInputSplit{
        if !isRoundValidChecker(gameResults){
            return false
        }
    }
    return true
}


