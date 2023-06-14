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
	// Rock = A = X = 1
	// Paper = B = Y = 2
	// Scissors = C = Z = 3
	// Win = 6, Draw = 3, Lose = 0
	// Score = Your choice + outcome of the game
	scoreMap := map[string]int{
		"AX": 1 + 3,
		"AY": 2 + 6,
		"AZ": 3 + 0,
		"BX": 1 + 0,
		"BY": 2 + 3,
		"BZ": 3 + 6,
		"CX": 1 + 6,
		"CY": 2 + 0,
		"CZ": 3 + 3,
	}

	lose := map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}

	draw := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	win := map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}

	strategyMap := map[string]map[string]string{
		"X": lose,
		"Y": draw,
		"Z": win,
	}

	score := 0
	for sc.Scan() {
		err = sc.Err()
		if err != nil {
			log.Panic(err)
		}

		line := sc.Text()
		game := strings.Fields(line)
		strategy := strategyMap[game[1]]
		// fmt.Println(game)
		// fmt.Println(strategy)
		score += scoreMap[game[0]+strategy[game[0]]]
	}
	fmt.Println(score)
}
