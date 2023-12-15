package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Numbers struct {
  number int
  index int
}

type Symbols struct {
  index int
  gears []Numbers
}

func Part2() int {
  file, err := os.Open("day.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  responseSum := 0
  symbolsLinesBuffer := make([][]Symbols, 2)
  numbersLinesBuffer := make([][]Numbers, 2)
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()

    symbolsIndexes := make([]Symbols, 0)
    for index, char := range line {
      if isGearSymbol(char) {
        symbolsIndexes = append(symbolsIndexes, Symbols{index, make([]Numbers, 0)})
      }
    }

    symbolsLinesBuffer = append(symbolsLinesBuffer[1:2], symbolsIndexes)

    numbersIndexes := extractNumbersIndexes(line)
    numbersLinesBuffer = append(numbersLinesBuffer[1:2], numbersIndexes)

    firstLine := numbersLinesBuffer[1]
    for _, symbolIndexes := range symbolsLinesBuffer {
      for sindex, symbolIndex := range symbolIndexes { 
        for lineIndex, lineNumber := range firstLine {
          if lineNumber.index - 1 <= symbolIndex.index && symbolIndex.index <= lineNumber.index + len(strconv.Itoa(lineNumber.number)) {
            symbolIndex.gears = append(symbolIndex.gears, lineNumber)
            symbolIndexes[sindex] = symbolIndex
            firstLine[lineIndex] = Numbers{index: -1, number: 1}
          }
        }
      }
    }

    secondLine := numbersLinesBuffer[0]
    for _, symbolIndexes := range symbolsLinesBuffer {
      for sindex, symbolIndex := range symbolIndexes { 
        for lineIndex, lineNumber := range secondLine {
          if lineNumber.index - 1 <= symbolIndex.index && symbolIndex.index <= lineNumber.index + len(strconv.Itoa(lineNumber.number)) {
            symbolIndex.gears = append(symbolIndex.gears, lineNumber)
            symbolIndexes[sindex] = symbolIndex
            secondLine[lineIndex] = Numbers{index: -1, number: 1}
          }
        }
      }
    }

    for _, symbolIndex := range symbolsLinesBuffer[0] {
      if len(symbolIndex.gears) == 2 {
        responseSum += symbolIndex.gears[0].number * symbolIndex.gears[1].number
      }
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return responseSum
}

func isGearSymbol(char rune) bool {
  return string(char) == "*"
}
