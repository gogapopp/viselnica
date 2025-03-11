package words

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
)

type Words struct {
	wordsFromTxt []string
}

func NewWords(filePath string) (Words, error) {
	words := Words{}

	wordsFromTxt, err := words.getWordsFromTxt(filePath)
	if err != nil {
		return Words{}, err
	}

	return Words{wordsFromTxt: wordsFromTxt}, nil
}

func (w *Words) getWordsFromTxt(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var words []string
	for scanner.Scan() {
		if scanner.Text() != "" {
			words = append(words, strings.ToLower(scanner.Text()))
		}
	}

	return words, nil
}

func (w *Words) GetRandomWord() string {
	r := rand.Intn(len(w.wordsFromTxt))
	return w.wordsFromTxt[r]
}
