package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strings"
)

func Part2() int {
  file, err := os.Open("day.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  responseSum := 0
  nextStratchsMap := make(map[int]int)
  lineCounter := 0
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    lineCounter++

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

    nextStratchsMap[lineCounter] += 1

    for i := 0; i < nextStratchsMap[lineCounter]; i++ {
      for w := 1; w <= winningNumbersCount; w++ {
        nextStratchsMap[lineCounter + w] += 1
      }
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  for _, value := range nextStratchsMap {
    responseSum += value
  }

  return responseSum
}
