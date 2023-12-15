package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Part1() int {
  file, err := os.Open("day.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  responseSum := 0
  symbolsLinesBuffer := make([][]int, 3)
  numbersLinesBuffer := make([][]struct{number, index int}, 3)
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()

    symbolsIndexes := make([]int, 0)
    for index, char := range line {
      if isSymbol(char) {
        symbolsIndexes = append(symbolsIndexes, index)
      }
    }

    symbolsLinesBuffer = append(symbolsLinesBuffer[1:3], symbolsIndexes)
    fmt.Println(symbolsLinesBuffer)

    numbersIndexes := extractNumbersIndexes(line)
    numbersLinesBuffer = append(numbersLinesBuffer[1:3], numbersIndexes)
    fmt.Println(numbersLinesBuffer)
    fmt.Println()

    if len(numbersLinesBuffer[1]) > 0 { //TODO: change verification to check always on new line or symbol detected
      actualLine := numbersLinesBuffer[1]
      for _, symbolIndexes := range symbolsLinesBuffer {
        for _, symbolIndex := range symbolIndexes { 
          for _, lineNumber := range actualLine {
            if lineNumber.index <= symbolIndex + 1 && symbolIndex <= lineNumber.index + len(strconv.Itoa(lineNumber.number)) + 1 {
              fmt.Println(lineNumber.number)
              responseSum += lineNumber.number
            }
          }
        }
      }
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return responseSum
}

func isSymbol(str rune) bool {
  return !unicode.IsNumber(str) && string(str) != "."
}

func extractNumbers(str string) []string {
  	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

  return strings.FieldsFunc(str, f)
}

func extractNumbersIndexes(line string) []struct{number, index int} {
  strs := extractNumbers(line)
  response := make([]struct{number, index int}, len(strs))

  for index, str := range strs {
    infos := struct{number, index int}{}
    infos.number, _ = strconv.Atoi(str)
    infos.index = strings.Index(line, str)

    response[index] = infos
  }

  return response
}
