package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

  f, err := os.Open("./input.txt")
  if err != nil {
    log.Panic(err)
  }
  
  sc := bufio.NewScanner(f)

  for sc.Scan() {
    err = sc.Err()
    if err != nil {
      log.Panic(err)
    }
    
    line := sc.Text()
    strategy := strings.Fields(line)
    fmt.Println(strategy)
  }

}
