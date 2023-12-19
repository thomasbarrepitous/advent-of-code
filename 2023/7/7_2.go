package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Cards []string
	Type  HandType
	Bid   int
}

func stringToHand(handStr string) Hand {
	handStrSplit := strings.Split(handStr, " ")
	handCards := strings.Split(handStrSplit[0], "")
	handType := calculateHandType(handCards)
	handBid, err := strconv.Atoi(handStrSplit[1])
	if err != nil {
		log.Fatal(err)
	}
	return Hand{Cards: handCards, Type: handType, Bid: handBid}
}

func calculateHandType(handCards []string) HandType {
	// Should have passed maps instead but this is just a CTF.
	if isFiveOfAKind(handCards) {
		return FiveOfAKind
	}
	if isFourOfAKind(handCards) {
		return FourOfAKind
	}
	if isFullHouse(handCards) {
		return FullHouse
	}
	if isThreeOfAKind(handCards) {
		return ThreeOfAKind
	}
	if isTwoPairs(handCards) {
		return TwoPairs
	}
	if isOnePair(handCards) {
		return OnePair
	}
	return HighCard
}

func isFiveOfAKind(handCards []string) bool {
	handMap := make(map[string]int)
	for _, card := range handCards {
		handMap[card]++
	}
	if handMap["J"] > 0 {
		// Handle case with Joker
		for _, match := range handMap {
			if match+handMap["J"] == 5 {
				return true
			}
		}
	}
	for _, match := range handMap {
		if match == 5 {
			return true
		}
	}
	return false
}

func isFourOfAKind(handCards []string) bool {
	handMap := make(map[string]int)
	for _, card := range handCards {
		handMap[card]++
	}
	if handMap["J"] > 0 {
		// Handle case with Joker
		for key, match := range handMap {
			if match+handMap["J"] == 4 && key != "J" {
				return true
			}
		}
	}
	for _, match := range handMap {
		if match == 4 {
			return true
		}
	}
	return false
}

func isFullHouse(handCards []string) bool {
	handMap := make(map[string]int)
	for _, card := range handCards {
		handMap[card]++
	}
	if handMap["J"] > 0 {
		nbPairs := 0
		// Handle case with Joker
		for _, match := range handMap {
			if match == 2 {
				nbPairs++
			}
		}
		if nbPairs == 2 {
			return true
		}
	}
	for _, match := range handMap {
		if match == 3 {
			for _, match2 := range handMap {
				if match2 == 2 {
					return true
				}
			}
		}
	}
	return false
}

func isThreeOfAKind(handCards []string) bool {
	handMap := make(map[string]int)
	for _, card := range handCards {
		handMap[card]++
	}
	if handMap["J"] > 0 {
		// Handle case with Joker
		for key, match := range handMap {
			if match+handMap["J"] == 3 && key != "J" {
				return true
			}
		}
	}
	for _, match := range handMap {
		if match == 3 {
			return true
		}
	}
	return false
}

func isTwoPairs(handCards []string) bool {
	handMap := make(map[string]int)
	for _, card := range handCards {
		handMap[card]++
	}
	if handMap["J"] > 0 {
		// Handle case with Joker
		for key, match := range handMap {
			if match == 2 && key != "J" {
				return true
			}
		}
	}
	nbPairs := 0
	for _, match := range handMap {
		if match == 2 {
			nbPairs++
		}
	}
	return nbPairs == 2
}

func isOnePair(handCards []string) bool {
	handMap := make(map[string]int)
	for _, card := range handCards {
		handMap[card]++
	}
	for key, match := range handMap {
		if match == 2 || (match == 1 && key == "J") {
			return true
		}
	}
	return false
}

func cardToInt(card string) int {
	if card == "J" {
		return 1
	}
	if card == "T" {
		return 10
	}
	if card == "Q" {
		return 11
	}
	if card == "K" {
		return 12
	}
	if card == "A" {
		return 13
	}
	parsedCard, err := strconv.Atoi(card)
	if err != nil {
		log.Fatal(err)
	}
	return parsedCard
}

func compareCard(card1 string, card2 string) int {
	// Follows slices.SortFunc implementation
	// Source :
	// https://cs.opensource.google/go/go/+/go1.21.5:src/slices/sort.go;l=26
	if cardToInt(card1) == cardToInt(card2) {
		return 0
	}
	if cardToInt(card1) > cardToInt(card2) {
		return 1
	}
	return -1
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

	hands := []Hand{}
	// Construct our sorted hands
	for _, line := range lines {
		hand := stringToHand(line)
		hands = append(hands, hand)
	}
	slices.SortFunc(hands, func(i, j Hand) int {
		if i.Type == j.Type {
			for index := 0; index < len(i.Cards); index++ {
				comparison := compareCard(i.Cards[index], j.Cards[index])
				if comparison != 0 {
					return comparison
				}
			}
		}
		return int(i.Type) - int(j.Type)
	})
	for _, hand := range hands {
		if hand.Type == TwoPairs {
			fmt.Println(hand)
		}
	}

	// Calculate flag
	for index := 0; index < len(hands); index++ {
		rank := index + 1
		flag += hands[index].Bid * rank
	}

	fmt.Println(flag)
}
