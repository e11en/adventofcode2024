package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	// Read the file content line by line
	scanner := bufio.NewScanner(file)
	content := []string{}
	total := 0

	for scanner.Scan() {
		text := scanner.Text()
		content = append(content, text)

		// Left to right
		total += strings.Count(text, "XMAS")

		// Right to left
		total += strings.Count(text, "SAMX")
	}

	// Get vertical
	for lineI, line := range content {
		for letterI, letter := range line {
			if lineI+3 >= len(content) {
				// No word possible
				continue
			}

			// Top to bottom
			if letter == 'X' {
				if content[lineI+1][letterI] == 'M' &&
					content[lineI+2][letterI] == 'A' &&
					content[lineI+3][letterI] == 'S' {
					total++
				}
			}
			// Reversed top to bottom
			if letter == 'S' {
				if content[lineI+1][letterI] == 'A' &&
					content[lineI+2][letterI] == 'M' &&
					content[lineI+3][letterI] == 'X' {
					total++
				}
			}
		}
	}

	// Get diagonal
	for lineI, line := range content {
		for letterI, letter := range line {
			if lineI+3 >= len(content) {
				// No word possible
				continue
			}

			if (letterI-3 >= 0) {
				// Top to bottom-left
				if letter == 'X' {
					if content[lineI+1][letterI-1] == 'M' &&
						content[lineI+2][letterI-2] == 'A' &&
						content[lineI+3][letterI-3] == 'S' {
						total++
					}
				}
				// Reverse top to bottom-left
				if letter == 'S' {
					if content[lineI+1][letterI-1] == 'A' &&
						content[lineI+2][letterI-2] == 'M' &&
						content[lineI+3][letterI-3] == 'X' {
						total++
					}
				}
			}
			

			if letterI+3 >= len(line) {
				// No word possible
				continue
			}

			// Top to bottom-right
			if letter == 'X' {
				if content[lineI+1][letterI+1] == 'M' &&
					content[lineI+2][letterI+2] == 'A' &&
					content[lineI+3][letterI+3] == 'S' {
					total++
				}
			}
			// Reverse top to bottom-right
			if letter == 'S' {
				if content[lineI+1][letterI+1] == 'A' &&
					content[lineI+2][letterI+2] == 'M' &&
					content[lineI+3][letterI+3] == 'X' {
					total++
				}
			}
		}
	}

	fmt.Println("Total", total)
}
