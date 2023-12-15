package main

import (
	"bufio"
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
  symbolsLinesBuffer := make([][]int, 2)
  numbersLinesBuffer := make([][]struct{number, index int}, 2)
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()

    symbolsIndexes := make([]int, 0)
    for index, char := range line {
      if isSymbol(char) {
        symbolsIndexes = append(symbolsIndexes, index)
      }
    }

    symbolsLinesBuffer = append(symbolsLinesBuffer[1:2], symbolsIndexes)

    numbersIndexes := extractNumbersIndexes(line)
    numbersLinesBuffer = append(numbersLinesBuffer[1:2], numbersIndexes)

    if (len(symbolsLinesBuffer[0]) > 0 || len(symbolsLinesBuffer[1]) > 0) && len(numbersLinesBuffer[0]) > 0 {
      actualLine := numbersLinesBuffer[0]
      for _, symbolIndexes := range symbolsLinesBuffer {
        for _, symbolIndex := range symbolIndexes { 
          for lineIndex, lineNumber := range actualLine {
            if lineNumber.index - 1 <= symbolIndex && symbolIndex <= lineNumber.index + len(strconv.Itoa(lineNumber.number)) {
              responseSum += lineNumber.number
              actualLine[lineIndex] = struct{number, index int}{}
            }
          }
        }
      }
    }

    if len(numbersLinesBuffer[1]) > 0 {
      actualLine := numbersLinesBuffer[1]
      for _, symbolIndexes := range symbolsLinesBuffer {
        for _, symbolIndex := range symbolIndexes { 
          for lineIndex, lineNumber := range actualLine {
            if lineNumber.index - 1 <= symbolIndex && symbolIndex <= lineNumber.index + len(strconv.Itoa(lineNumber.number)) {
              responseSum += lineNumber.number
              actualLine[lineIndex] = struct{number, index int}{}
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
  return !unicode.IsDigit(str) && string(str) != "."
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
    line = strings.Replace(line, str, strings.Repeat(".", len(str)) ,1)

    response[index] = infos
  }

  return response
}
