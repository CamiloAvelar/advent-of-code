package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day1_2() int {
  strIntMap := map[string]int{
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
  }

  file, err := os.Open("day1-1.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  var result int

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    var first, last = struct{ index, value int }{-1, -1}, struct{ index, value int }{-1, -1}

    for strChar, intChar := range strIntMap {
      if(strings.Contains(line, strChar)) {
        if first.index > strings.Index(line, strChar) || first.index == -1 {
          first.index = strings.Index(line, strChar)
          first.value = intChar
        }

        if last.index < strings.LastIndex(line, strChar) {
          last.index = strings.LastIndex(line, strChar)
          last.value = intChar
        }
      }
    }

    for charIndex, char := range line {
      strChar := string(char)

      if intChar, err := strconv.Atoi(strChar); err == nil {
        if first.index > charIndex || first.index == -1 {
          first.index = charIndex
          first.value = intChar
        }

        if last.index < charIndex {
          last.index = charIndex
          last.value = intChar
        }
      }
    }
    
    lineResult := fmt.Sprintf("%d%d", first.value, last.value)
    intLineResult, _ := strconv.Atoi(lineResult)

    result += intLineResult
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return result
}
