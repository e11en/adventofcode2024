package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	// Read the file content line by line
	scanner := bufio.NewScanner(file)
	content := ""
	for scanner.Scan() {
		content += scanner.Text()
	}

	total := 0
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(content, -1)
	for _, match := range matches {
		numbers := strings.Split(match[4:len(match)-1], ",")
		firstNumber, _ := strconv.Atoi(numbers[0])
		secondNumber, _ := strconv.Atoi(numbers[1])
		total += firstNumber * secondNumber
	}

	fmt.Println("Total", total)
}
