package gamesession

import (
	"bufio"
	"fmt"
	"strings"
	"viselnica/internal/hangman"
	"viselnica/internal/words"
)

type gameSession struct {
	allWords           words.Words
	userMaxAttemptions int
	hangman            hangman.Hangman

	userMistakes        int
	gameWord            string
	userWordCondition   []string
	usedLetters         map[string]struct{}
	numOfGuessedLetters int
}

func NewGameSession(allWords words.Words, userAttemptions int, hangman hangman.Hangman) gameSession {
	return gameSession{
		allWords:           allWords,
		userMaxAttemptions: userAttemptions,
		hangman:            hangman,
	}
}

func (gs *gameSession) Reset() {
	gs.userMistakes = 0
	gs.gameWord = gs.allWords.GetRandomWord()
	gs.userWordCondition = make([]string, len(gs.gameWord))
	gs.usedLetters = make(map[string]struct{})
	gs.numOfGuessedLetters = 0
}

func (gs *gameSession) Start(userInput *bufio.Scanner) {
	fmt.Printf("Попробуй угадать секретное слово из %d букв\n", len([]rune(gs.gameWord)))
	fmt.Println("На вход требуется только одна буква")

	for userInput.Scan() {
		input := userInput.Text()

		if !gs.isValidUserInput([]rune(input)) {
			fmt.Println("Неправильный ввод, требуется только один символ кирилицы")
			continue
		}

		if gs.isLetterUsed(input) {
			fmt.Println("Вы уже использовали эту букву")
			continue
		}

		if gs.checkSymbol(input) {
			fmt.Println("Вы угадали букву!")

			if len([]rune(gs.gameWord)) == gs.numOfGuessedLetters {
				fmt.Println("Вы выиграли!")
				fmt.Println(gs.userWordCondition)
				break
			}

			fmt.Printf("Осталось попыток %d\n", gs.userMaxAttemptions-gs.userMistakes)
		} else {
			fmt.Println("Вы не угадали букву!")

			hangManState := gs.userMistakes - 1
			gs.hangman.Draw(hangManState)

			fmt.Printf("Осталось попыток %d\n", gs.userMaxAttemptions-gs.userMistakes)

			if gs.userMistakes == gs.userMaxAttemptions {
				fmt.Println("Вы проиграли!")
				fmt.Println(gs.userWordCondition)
				break
			}
		}
		fmt.Println(gs.userWordCondition)
		fmt.Println("Введите букву: ")
		fmt.Println("")
	}
}

func (gs *gameSession) isValidUserInput(input []rune) bool {
	if len(input) > 1 || len(input) == 0 {
		return false
	}

	str := string(input)

	m := map[string]struct{}{
		"0": {},
		"1": {},
		"2": {},
		"3": {},
		"4": {},
		"5": {},
		"6": {},
		"7": {},
		"8": {},
		"9": {},
	}
	if _, ok := m[str]; ok {
		return false
	}

	return true
}

func (gs *gameSession) isLetterUsed(str string) bool {
	str = strings.ToLower(str)

	if _, ok := gs.usedLetters[str]; ok {
		return true
	}

	gs.usedLetters[str] = struct{}{}

	return false
}

func (gs *gameSession) checkSymbol(symbol string) bool {
	found := false

	symbol = strings.ToLower(symbol)

	s := []rune(symbol)[0]

	for i, ch := range gs.gameWord {
		if ch == s {
			gs.userWordCondition[i] = string(ch)
			gs.increaseNumOfGuessedLetters()
			found = true
		}
	}

	if !found {
		gs.increaseUserMistakes()
	}

	return found
}

func (gs *gameSession) increaseNumOfGuessedLetters() {
	gs.numOfGuessedLetters++
}

func (gs *gameSession) increaseUserMistakes() {
	gs.userMistakes++
}
