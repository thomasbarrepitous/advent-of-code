package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cardToUsableArrays(card string) (winningNumbers []int, cardNumbers []int) {
	cardSplit := strings.Split(card, ":")
	cardInputSplit := strings.Split(cardSplit[1], "|")
	winningNumbersStr := strings.Split(cardInputSplit[0], " ")
	cardNumbersStr := strings.Split(cardInputSplit[1], " ")
	// Parse to int winning numbers
	for _, number := range winningNumbersStr {
		winningNumber, err := strconv.Atoi(number)
		if err != nil {
			continue
		}
		winningNumbers = append(winningNumbers, winningNumber)
	}
	// Parse to int card numbers
	for _, number := range cardNumbersStr {
		cardNumber, err := strconv.Atoi(number)
		if err != nil {
			continue
		}
		cardNumbers = append(cardNumbers, cardNumber)
	}
	return
}

func getCardPoints(card string) (points int) {
	matches := 0
	winningNumbers, cardNumbers := cardToUsableArrays(card)
	for _, winningNumber := range winningNumbers {
		for _, cardNumber := range cardNumbers {
			if winningNumber == cardNumber {
				matches++
				break
			}
		}
	}
	fmt.Println("Matches : ", matches)
	for i := 0; i < matches; i++ {
		if i == 0 {
			points = 1
		} else {
			points = points * 2
		}
		fmt.Println("Points : ", points)
	}
	return
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
	for _, line := range lines {
		fmt.Println(line)
		flag += getCardPoints(line)
		fmt.Println("Flag : ", flag)
	}
	fmt.Println(flag)
}
