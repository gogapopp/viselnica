package gamesession

import (
	"bufio"
	"fmt"
	"unicode"
	"viselnica/internal/hangman"
	"viselnica/internal/words"
)

type GameSession struct {
	allWords           words.Words
	userMaxAttempts    int
	hangmanAsciiStages hangman.Hangman

	userMistakes        int
	gameWord            string
	userWordCondition   []rune
	usedLetters         map[rune]struct{}
	numOfGuessedLetters int
}

func NewGameSession(allWords words.Words, userAttemptions int, hangmanAscii hangman.Hangman) GameSession {
	return GameSession{
		allWords:           allWords,
		userMaxAttempts:    userAttemptions,
		hangmanAsciiStages: hangmanAscii,
	}
}

func (gs *GameSession) Reset() {
	gs.userMistakes = 0
	gs.gameWord = gs.allWords.GetRandomWord()
	gs.userWordCondition = make([]rune, len([]rune(gs.gameWord)))
	for i := range gs.userWordCondition {
		gs.userWordCondition[i] = '_'
	}
	gs.usedLetters = make(map[rune]struct{})
	gs.numOfGuessedLetters = 0
}

func (gs *GameSession) Start(userInput *bufio.Scanner) {
	fmt.Printf("Попробуй угадать секретное слово из %d букв\n", len([]rune(gs.gameWord)))
	fmt.Println("На вход требуется только одна буква")

	for userInput.Scan() {
		textUserInput := userInput.Text()

		if gs.handleUserInput(textUserInput) {
			break
		}

		fmt.Println(string(gs.userWordCondition))
		fmt.Println("Введите букву: ")
		fmt.Println("")
	}
}

func (gs *GameSession) handleUserInput(textUserInput string) bool {
	if !gs.isValidUserInput([]rune(textUserInput)) {
		fmt.Println("Неправильный ввод, требуется только один символ кирилицы")
		return false
	}

	runeUserInput := unicode.ToLower([]rune(textUserInput)[0])

	if gs.isLetterUsed(runeUserInput) {
		fmt.Println("Вы уже использовали эту букву")
		return false
	}

	if gs.processUserGuess(runeUserInput) {
		fmt.Println("Вы угадали букву!")

		if len([]rune(gs.gameWord)) == gs.numOfGuessedLetters {
			fmt.Println("Вы выиграли!")
			fmt.Println(string(gs.userWordCondition))
			return true
		}
	} else {
		fmt.Println("Вы не угадали букву!")
		hangManState := gs.userMistakes - 1
		gs.hangmanAsciiStages.Draw(hangManState)

		if gs.userMistakes == gs.userMaxAttempts {
			fmt.Println("Вы проиграли!")
			fmt.Println(string(gs.userWordCondition))
			return true
		}
	}

	return false
}

func (gs *GameSession) isValidUserInput(input []rune) bool {
	if len(input) > 1 || len(input) == 0 {
		return false
	}

	isLetter := unicode.IsLetter(rune(input[0]))
	if !isLetter {
		return false
	}

	return true
}

func (gs *GameSession) isLetterUsed(input rune) bool {
	if _, ok := gs.usedLetters[input]; ok {
		return true
	}

	gs.usedLetters[input] = struct{}{}

	return false
}

func (gs *GameSession) processUserGuess(symbol rune) bool {
	found := false

	gameWordRunes := []rune(gs.gameWord)

	for i, ch := range gameWordRunes {
		if ch == symbol {
			gs.userWordCondition[i] = ch
			gs.increaseNumOfGuessedLetters()
			found = true
		}
	}

	if !found {
		gs.increaseUserMistakes()
	}

	return found
}

func (gs *GameSession) increaseNumOfGuessedLetters() {
	gs.numOfGuessedLetters++
}

func (gs *GameSession) increaseUserMistakes() {
	gs.userMistakes++
}
