package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strings"
)

func Part1() int {
  file, err := os.Open("day.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  responseSum := 0
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()

    firstPartWithPrefix := strings.Split(line, "|")[0]
    winningPart := strings.Split(firstPartWithPrefix, ":")[1]
    myNumbersPart := strings.Split(line, "|")[1]

    winningNumbersStr := strings.Split(winningPart, " ")
    myNumbersStr := strings.Split(myNumbersPart, " ")

    winningNumbersCount := 0;

    for _, myNumber := range myNumbersStr {
      if myNumber == "" {
        continue
      }

      if slices.Contains(winningNumbersStr, myNumber) {
        winningNumbersCount++
      }
    }

    responseSum += calculateResult(winningNumbersCount)
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return responseSum
}

func calculateResult (count int) int {
  result := 0

  for i := 0; i < count; i++ {
    if i == 0 {
      result = 1
    } else {
      result = result * 2
    }
  }

  return result
}
