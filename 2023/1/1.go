package main


import (
    "bufio"
    "os"
    "fmt"
)


func reverse(str string) (result string) { 
    for _, v := range str { 
        result = string(v) + result 
    } 
    return
}

func findNumberInString(s string) int {
    var number int
    for _, c := range s {
        if c >= '0' && c <= '9' {
            number = int(c - '0')
            break
        }
    }
    return number
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
		// fmt.Println(line)
        firstNb := findNumberInString(line)
        lastNb := findNumberInString(reverse(line))
        res := firstNb*10 + lastNb
        // fmt.Println(res)
        total += res
	}
    fmt.Println(total)
}

