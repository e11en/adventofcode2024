package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Abs(x int) int {
  if x >= 0 {
    return x
  } else {
    return -x
  }
}

func IsSafeSequence(sequence []int) bool {
  // Part 1
  if len(sequence) < 2 { return true }
  
  isAscending := sequence[1] - sequence[0] > 0
  pastNumber := sequence[0]

  for _, currentNumber := range sequence[1:] {
    grad := currentNumber - pastNumber

    if grad == 0 {
      return false // neither an increase or a decrease
    } else if (grad > 0 && !isAscending) || (grad < 0 && isAscending) { // ascending or desecnding case
      return false // wrong direction
    }
    if !(Abs(grad) >= 1 && Abs(grad) <= 3) {
      return false // an increase / derease of more than 3
    }

    pastNumber = currentNumber
  }

  return true
}

func SliceAt(slice []int, index int) []int {
  var s []int
	
	s = append(s, slice[:index]...)
	s = append(s, slice[index+1:]...)
	
	return s
}

func IsSafeSequenceDamped(sequence []int) bool {
  // Part 2
  if len(sequence) < 2 { return true }

  isSafe := IsSafeSequence(sequence)

  if !isSafe {
    isNewSafe := false
    for i := 0 ; i < len(sequence); i++ {
      newSlice := SliceAt(sequence, i)

      if IsSafeSequence(newSlice) {
        isNewSafe = true
        break
      }
    }

    if !isNewSafe {
      return false
    }
  }

  return true
}

func DetectUnusualSequence(rawInput string) int {
  splitByLine := strings.Split(rawInput, "\n")

  safeReports := 0

  for _, line := range splitByLine {
    splitBySpace := strings.Split(line, " ")

    numLevels := make([]int, len(splitBySpace))

    for i := 0; i < len(splitBySpace); i++ {
      intValue, err := strconv.Atoi(splitBySpace[i])
      check(err)
      numLevels[i] = intValue
    }

    // if IsSafeSequence(numLevels) {
    if IsSafeSequenceDamped(numLevels) {
      safeReports += 1
    }
  }

  return safeReports
}

func main() {
	b1, err := os.ReadFile("levels.txt")
  check(err)

	data := string(b1)

	fmt.Println(DetectUnusualSequence(data))
}