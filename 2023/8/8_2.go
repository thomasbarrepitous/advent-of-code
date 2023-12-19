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

func areNodesAllAtFinish(currentNodes []string) bool {
	for _, currentNode := range currentNodes {
		if string(currentNode[2]) != "Z" {
			return false
		}
	}
	return true
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
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

	currentNodes := make([]string, 0)
	for key := range nodeMap {
		if string(key[2]) == "A" {
			currentNodes = append(currentNodes, key)
		}
	}

	steps := make([]int, len(currentNodes))
nodeLoop:
	for idx := range currentNodes {
		for {
			for _, instruction := range instructions {
				if string(currentNodes[idx][2]) == "Z" {
					continue nodeLoop
				}
				steps[idx]++
				// Note that we fetching the value and then modifying it
				// wouldn't work as it returns a copy of the value.
				// Thus we need to use the index to modify the slice.
				fmt.Println(currentNodes)
				if instruction == 'L' {
					currentNodes[idx] = nodeMap[currentNodes[idx]].leftElement
				}
				if instruction == 'R' {
					currentNodes[idx] = nodeMap[currentNodes[idx]].rightElement
				}
			}
		}
	}
	// Thanks to whoever wrote those LCM and GCD functions for me
	fmt.Println(LCM(steps[0], steps[1], steps[2:]...))
	fmt.Println(flag)
}
