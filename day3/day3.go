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
	updatedMul, _ := regexp.Compile(`(mul\([0-9]+,[0-9]+\)|do(n't)??\(\))`)
	enabled := true
	findNum, _ := regexp.Compile(`([0-9]+)`)
	total2 := 0

	for scanner.Scan() {
		content += scanner.Text()

		text := scanner.Text()
		matches := updatedMul.FindAllString(text, -1)

		for _, match := range matches {
			if match == "do()" {
				enabled = true
				continue
			} else if match == "don't()" {
				enabled = false
				continue
			}
			if enabled {
				nums := findNum.FindAllString(match, -1)
				first, _ := strconv.Atoi(nums[0])
				second, _ := strconv.Atoi(nums[1])
				total2 += first * second
			}
		}
	}

	total1 := 0
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(content, -1)
	for _, match := range matches {
		numbers := strings.Split(match[4:len(match)-1], ",")
		firstNumber, _ := strconv.Atoi(numbers[0])
		secondNumber, _ := strconv.Atoi(numbers[1])
		total1 += firstNumber * secondNumber
	}

	fmt.Println("Total 1", total1)
	fmt.Println("Total 2", total2)
}
