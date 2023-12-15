package main

import (
	"bufio"
	"log"
	"os"
)

func Part2() int {
  file, err := os.Open("day.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  responseSum := 0
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return responseSum
}
