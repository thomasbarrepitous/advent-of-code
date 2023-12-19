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
	for _, card := range handCards {
		if card != handCards[0] {
			return false
		}
	}
	return true
}

func isFourOfAKind(handCards []string) bool {
	match := 0
	cardMatch := handCards[0]
	copyHand := make([]string, len(handCards))
	copy(copyHand, handCards)
	slices.Sort(copyHand)
	for _, card := range copyHand {
		if card == cardMatch {
			match++
		} else {
			if match == 4 {
				return true
			}
			match = 1
			cardMatch = card
		}
	}
	return match == 4
}

func isFullHouse(handCards []string) bool {
	matchesMap := make(map[string]int)
	for _, card := range handCards {
		matchesMap[card]++
	}
	for _, match := range matchesMap {
		if match == 3 {
			for _, matchSecond := range matchesMap {
				if matchSecond == 2 {
					return true
				}
			}
		}
	}
	return false
}

func isThreeOfAKind(handCards []string) bool {
	match := 0
	cardMatch := handCards[0]
	copyHand := make([]string, len(handCards))
	copy(copyHand, handCards)
	slices.Sort(copyHand)
	for _, card := range copyHand {
		if card == cardMatch {
			match++
		} else {
			if match == 3 {
				return true
			}
			match = 1
			cardMatch = card
		}
	}
	return match == 3
}

func isTwoPairs(handCards []string) bool {
	matchesMap := make(map[string]int)
	for _, card := range handCards {
		matchesMap[card]++
	}
	pairs := 0
	for _, match := range matchesMap {
		if match == 2 {
			pairs++
		}
	}
	return pairs == 2
}

func isOnePair(handCards []string) bool {
	match := 0
	cardMatch := handCards[0]
	copyHand := make([]string, len(handCards))
	copy(copyHand, handCards)
	slices.Sort(copyHand)
	for _, card := range copyHand {
		if card == cardMatch {
			match++
		} else {
			if match == 2 {
				return true
			}
			cardMatch = card
			match = 1
		}
	}
	return match == 2
}

func cardToInt(card string) int {
	if card == "T" {
		return 10
	}
	if card == "J" {
		return 11
	}
	if card == "Q" {
		return 12
	}
	if card == "K" {
		return 13
	}
	if card == "A" {
		return 14
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

	// Calculate flag
	for index := 0; index < len(hands); index++ {
		rank := index + 1
		flag += hands[index].Bid * rank
	}

	fmt.Println(flag)
}
