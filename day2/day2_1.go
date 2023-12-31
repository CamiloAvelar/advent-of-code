package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var MaxColors = map[string]int{
  "blue": 14,
  "green": 13,
  "red": 12,
}

func Day2_1() int {
  file, err := os.Open("day2.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  validGamesSum := 0
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    splitLine := strings.Split(line, ":")

    gameId := strings.Split(splitLine[0], " ")[1]
    intGameId, _ := strconv.Atoi(gameId)

    gamesString := strings.Split(splitLine[1], ";")

    playIsValid := true
    for _, gameString := range gamesString {
      playsString := strings.Split(gameString, ",")

      for _, playString := range playsString {
        play := strings.TrimSpace(playString)
        splitedPlay := strings.Split(play, " ")

        color := splitedPlay[1]
        number, _ := strconv.Atoi(splitedPlay[0])

        if number > MaxColors[color] {
          playIsValid = false
        }
      }
    }

    if playIsValid {
      validGamesSum += intGameId
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return validGamesSum
}
