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
        powerSetOfCube := getPowerMinSetOfCube(lineSplit[1])
        flag += powerSetOfCube
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

func getPowerMinSetOfCube(gameInput string) int{
    red, green, blue:= getMinimumSetOfCube(gameInput)
    return red * green * blue
}

func getRoundColourValues(round string) (red int,green int,blue int){
    var err error
    red, green, blue = 0, 0, 0
    colourSplit := strings.Split(round, ",")
    for _, cube:= range colourSplit {
        colour := colourChecker(cube)
        cubeSplit:= strings.Split(cube, colour)
        if colour == "red" {
            red, err = strconv.Atoi(strings.TrimSpace(cubeSplit[0]))
            if err != nil {
                log.Fatal(err)
            }
        }
        if colour == "green" {
            green, err = strconv.Atoi(strings.TrimSpace(cubeSplit[0]))
            if err != nil {
                log.Fatal(err)
            }
        }
        if colour == "blue" {
            blue, err  = strconv.Atoi(strings.TrimSpace(cubeSplit[0]))
            if err != nil {
                log.Fatal(err)
            }
        }
    }
    return
}

func getMinimumSetOfCube(gameInput string) (redMax int, greenMax int, blueMax int){
    redMax, greenMax, blueMax = 0, 0, 0
    fmt.Println(gameInput)
    // Delimiter round
    roundSplit := strings.Split(gameInput, ";")
    // For each round get max cube
    for _, round := range roundSplit {
        red, green, blue := getRoundColourValues(round)
        fmt.Println(red, green, blue)
        if red > redMax {
            redMax = red
        }
        if green > greenMax {
            greenMax = green
        }
        if blue > blueMax {
            blueMax = blue
        }
    }
    return 
}
