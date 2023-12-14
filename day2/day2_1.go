package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day2_1() int {
  file, err := os.Open("day2.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()

    gameId := strings.Split(line, " ")[1]
    intGameId, _ := strconv.Atoi(strings.Trim(gameId, ":"))

    gamesString := strings.Split(line, ";")

    for _, gameString := range gamesString {
      fmt.Println(gameString)
      playsString := strings.Split(gameString, ",")
      fmt.Println(playsString)
    }
    fmt.Println(intGameId)
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return 0
}
