package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Map map[string]map[string]string

func Part1() int {
  file, err := os.Open("day.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  responseSum := 0
  scanner := bufio.NewScanner(file)
  seedsMap := make(Map)
  var currentOrigin string
  var currentDestination string
  for scanner.Scan() {
    line := scanner.Text()

    if strings.Contains(line, "seeds") {
      seedsPart := strings.Split(line, ":")[1]

      seeds := strings.Split(seedsPart, " ")
      for _, seed := range seeds {
        if seed == " " || seed == "" {
          continue
        }

        seedsMap[seed] = make(map[string]string)
        seedsMap[seed]["seed"] = seed
      }
      continue
    }

    if strings.Contains(line, "map") {
      mapPart := strings.Split(line, " ")[0]
      originDestination := strings.Split(mapPart, "-to-")
      currentOrigin = originDestination[0]
      currentDestination = originDestination[1]
      continue
    }

    if line == "" {
      continue
    }

    for seed, _ := range seedsMap {
      if seedsMap[seed][currentOrigin] != seedsMap[seed][currentDestination] && seedsMap[seed][currentOrigin] != "" && seedsMap[seed][currentDestination] != "" {
        continue
      }

      seedsMap[seed][currentDestination] = originToDestination(line, seedsMap[seed][currentOrigin])
    }
  }

  var lowestValue string
  for _, seedMap := range seedsMap {
    for destination, value := range seedMap {
      if destination == "location" {
        if lowestValue == "" {
          lowestValue = value
        } else {
          a, _ := strconv.Atoi(value)
          b, _ := strconv.Atoi(lowestValue)
          if a < b {
            lowestValue = value
          }
        }
      }
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  responseSum, _ = strconv.Atoi(lowestValue)
  return responseSum
}

func originToDestination(line, origin string) string {
  lineMappings := strings.Split(line, " ")
  initialDestination, _ := strconv.Atoi(lineMappings[0])
  initialOrigin, _ := strconv.Atoi(lineMappings[1])
  offset, _ := strconv.Atoi(lineMappings[2])
  intOrigin, _ := strconv.Atoi(origin)

  if intOrigin >= initialOrigin && intOrigin <= initialOrigin + offset {
    return strconv.Itoa(initialDestination + (intOrigin - initialOrigin))
  }

  return origin
}
