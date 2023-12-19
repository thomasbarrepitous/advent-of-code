package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	rightElement string
	leftElement  string
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

	instructions := lines[0]
	nodeMap := make(map[string]Node)
	// Construct input
	for _, line := range lines[2:] {
		lineSplit := strings.Split(line, "=")
		key := strings.TrimSpace(lineSplit[0])
		elements := strings.Split(lineSplit[1], ",")
		leftElement := strings.TrimSpace(elements[0][2:])
		rightElement := strings.TrimSpace(elements[1][:len(elements[1])-1])
		node := Node{leftElement: leftElement, rightElement: rightElement}
		nodeMap[key] = node
	}

	currentNode := "AAA"
outerLoop:
	for {
		for _, instruction := range instructions {
			if currentNode == "ZZZ" {
				break outerLoop
			}
			flag++
			if instruction == 'L' {
				currentNode = nodeMap[currentNode].leftElement
			}
			if instruction == 'R' {
				currentNode = nodeMap[currentNode].rightElement
			}
		}
	}
	fmt.Println(flag)
}
