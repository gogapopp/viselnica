package main

import (
	"bufio"
	"fmt"
)

type gameSession struct {
	allWords           words
	userMaxAttemptions int

	userMistakes        int
	gameWord            string
	userWordCondition   []string
	usedLetters         map[string]struct{}
	numOfGuessedLetters int
}

func NewGameSession(allWords words, userAttemptions int) gameSession {
	return gameSession{
		allWords:           allWords,
		userMaxAttemptions: userAttemptions,
	}
}

func (gs *gameSession) Start(userInput *bufio.Scanner) {
	if len([]rune(gs.gameWord)) < 5 {
		fmt.Printf("Попробуй угадать секретное слово из %d букв\n", len([]rune(gs.gameWord)))
		fmt.Println("Так как слово меньше или равно четырём символам, то играем без подсказок")
		fmt.Println("На вход требуется только одна буква")

	} else {
		fmt.Printf("Попробуй угадать секретное слово из %d букв\n", len([]rune(gs.gameWord)))
		fmt.Println("На вход требуется только одна буква")
	}

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
	if _, ok := gs.usedLetters[str]; ok {
		return true
	}

	gs.usedLetters[str] = struct{}{}

	return false
}

func (gs *gameSession) checkSymbol(symbol string) bool {
	found := false

	s := []rune(symbol)[0]

	for i, ch := range gs.gameWord {
		if ch == s {
			gs.userWordCondition[i] = string(ch)
			gs.numOfGuessedLetters++
			found = true
		}
	}

	if !found {
		gs.userMistakes++
	}

	return found
}

func (gs *gameSession) Reset() {
	gs.userMistakes = 0
	gs.gameWord = gs.allWords.getRandomWord()
	gs.userWordCondition = make([]string, len(gs.gameWord))
	gs.usedLetters = make(map[string]struct{})
	gs.numOfGuessedLetters = 0
}
