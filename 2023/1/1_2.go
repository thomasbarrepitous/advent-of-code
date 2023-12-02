package main


import (
    "bufio"
    "os"
    "fmt"
)

var lettersToInt = map[string]int{
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
}

var numbersInLetters = []string{
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
}


func reverse(str string) (result string) { 
    for _, v := range str { 
        result = string(v) + result 
    } 
    return
}

func findNumberInString(s string, isStringReversed bool) (number int) {
    for idx, c := range s {
        // Check if the character is a number
        if c >= '0' && c <= '9' {
            number = int(c - '0')
            break
        }
        // Check if surrounding characters are matching numbers
        if idx > 0 && idx < len(s)-1 {
            for _, numberInLetter := range numbersInLetters {
                if isStringReversed {
                    val, ok := isStringNumberCheck(s, idx, reverse(numberInLetter))
                    if ok {
                        return lettersToInt[reverse(val)]
                    }
                }
                val, ok := isStringNumberCheck(s, idx, numberInLetter)
                if ok {
                    return lettersToInt[val]
                }
            }
        }
    }
    return
}
// s : the current line of the file
// idx : the index of s
// v : the checked number in letters
func isStringNumberCheck(s string, idx int, v string) (string, bool){
    // Check if surrounding characters are matching
    if string(s[idx-1]) == string(v[0]) && string(s[idx]) == string(v[1]) && string(s[idx+1]) == string(v[2]) {
        // Check if the string is long enough to contain the number
        if len(s[idx:]) < len(v) {
            return "", false
        }
        subString := string(s[idx-1:idx+len(v)-1])
        for k, letter := range v {
            if string(subString[k]) != string(letter) {
                return "", false
            }
        }
        return v, true
    }
    return "", false
}

func main() {
    // Specify the path to your text file
	filePath := "input_1.txt"

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

    var total int
	// Print the lines or perform other operations as needed
	for _, line := range lines {
		fmt.Println(line)
        firstNb := findNumberInString(line, false)
        lastNb := findNumberInString(reverse(line), true)
        res := firstNb*10 + lastNb
        fmt.Println(res)
        total += res
	}
    fmt.Println(total)
}

