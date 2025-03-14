package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	gamesession "viselnica/internal/game_session"
	"viselnica/internal/hangman"
	"viselnica/internal/words"
)

const (
	// т.к. 5 ASCII вариантов висельника
	userMaxAttemptions = 5
	wordsPath          = "./words.txt"
)

func main() {
	words, err := words.NewWords(wordsPath)
	if err != nil {
		log.Fatal(err)
	}

	hangmanAscii := hangman.NewHangman()

	gameSession := gamesession.NewGameSession(words, userMaxAttemptions, hangmanAscii)

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
