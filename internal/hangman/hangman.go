package hangman

import "fmt"

type Hangman struct {
	hangman []string
}

func NewHangman() Hangman {
	return Hangman{
		hangman: hangmanASCII,
	}
}

func (hm Hangman) Draw(state int) {
	fmt.Println(hm.hangman[state])
}

var hangmanASCII = []string{
	`+---+
  |   |
      |
      |
      |
      |
=========`,
	`+---+
  |   |
  O   |
      |
      |
      |
=========`,
	`+---+
  |   |
  O   |
 /|   |
      |
      |
=========`,
	`+---+
  |   |
  O   |
 /|\  |
      |
      |
=========`,
	`+---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========`,
}
