package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day2_2() int {
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

    gamesString := strings.Split(splitLine[1], ";")

    gameMinimumSet := map[string]int{
      "blue": 1,
      "green": 1,
      "red": 1,
    }

    for _, gameString := range gamesString {
      playsString := strings.Split(gameString, ",")

      for _, playString := range playsString {
        play := strings.TrimSpace(playString)
        splitedPlay := strings.Split(play, " ")

        color := splitedPlay[1]
        number, _ := strconv.Atoi(splitedPlay[0])

        if number > gameMinimumSet[color] {
          gameMinimumSet[color] = number
        }
      }
    }

    validGamesSum += gameMinimumSet["blue"] * gameMinimumSet["green"] * gameMinimumSet["red"]
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return validGamesSum
}
