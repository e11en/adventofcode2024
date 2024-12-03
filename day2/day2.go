package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the input file
	file, err := os.Open("levels.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file content line by line
	scanner := bufio.NewScanner(file)
	var content []string
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	counter := 0

	// Process each report
	for _, report := range content {
		// Parse the values in the report
		strValues := strings.Fields(report)
		var values []int
		for _, v := range strValues {
			num, _ := strconv.Atoi(v)
			values = append(values, num)
		}

		// Initialize safepos and safeneg sets
		safepos := map[int]struct{}{1: {}, 2: {}, 3: {}}
		safeneg := map[int]struct{}{-1: {}, -2: {}, -3: {}}

		// Check the differences between consecutive values
		for i := 1; i < len(values); i++ {
			diff := values[i] - values[i-1]
			safepos[diff] = struct{}{}
			safeneg[diff] = struct{}{}
		}

		// If either set has exactly 3 unique elements, increment the answer
		if len(safepos) == 3 || len(safeneg) == 3 {
			counter++
		}
	}

	fmt.Println(counter)
}
