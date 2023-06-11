package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
  fmt.Println("Hello Advent of Code 2022!")


  f, err := os.Open("./input.txt")
  if err != nil {
    log.Panic(err)
  }
  
  sc := bufio.NewScanner(f)

  maxCal := 0
  maxCal2 := 0
  maxCal3 := 0
  totalCal := 0

  for sc.Scan() {
    err := sc.Err()
    if err != nil {
      log.Panic(err)
    }
    
    line := sc.Bytes()
    if len(line) == 0 { 
      totalCal = 0
      continue
    }
    
    cal, err := strconv.Atoi(string(line))
    if err != nil {
      log.Panic(err)
    }
    totalCal += cal

    if totalCal > maxCal {
      maxCal3 = maxCal2
      maxCal2 = maxCal
      maxCal = totalCal
    } else if totalCal > maxCal2 {
      maxCal3 = maxCal2
      maxCal2 = totalCal
    } else if totalCal > maxCal3 {
      maxCal3 = totalCal
    }

    fmt.Println(maxCal, maxCal2, maxCal3)
    
    // fmt.Println(fmt.Sprintf("Elf %d, calories: %v", elf, totalCal))
  }

  fmt.Println(fmt.Sprintf("Max cal: %d, %d, %d, sum: %d", maxCal, maxCal2, maxCal3, maxCal + maxCal2 + maxCal3))

}

