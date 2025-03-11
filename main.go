package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	userMaxAttemptions = 3
)

func main() {
	words, err := NewWords("words.txt")
	if err != nil {
		log.Fatal(err)
	}

	gameSession := NewGameSession(words, userMaxAttemptions)

	for {
		fmt.Println("Начать новую игру - 0, выйти из игры - 1")

		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			if scanner.Text() == "1" {
				os.Exit(0)
			}

			if scanner.Text() == "0" {
				gameSession.Reset()
				gameSession.Start(scanner)
				break
			}
		}
	}
}
