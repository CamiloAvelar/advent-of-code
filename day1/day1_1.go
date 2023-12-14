package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day1_1() int {
  file, err := os.Open("day1-1.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  var result int

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    var first, last = -1, -1

    for _, char := range line {
      strChar := string(char)

      if intChar, err := strconv.Atoi(strChar); err == nil {
        if first == -1 {
          first = intChar
        }

        last = intChar
      }
    }
    
    lineResult := fmt.Sprintf("%d%d", first, last)
    intLineResult, _ := strconv.Atoi(lineResult)

    result += intLineResult
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return result
}
