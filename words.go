package main

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
)

type words struct {
	wordsFromTxt []string
}

func NewWords(filePath string) (words, error) {
	wordsFromTxt, err := getWordsFromTxt(filePath)
	if err != nil {
		return words{}, err
	}

	return words{wordsFromTxt: wordsFromTxt}, nil
}

func getWordsFromTxt(filePath string) ([]string, error) {
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

func (w words) getRandomWord() string {
	r := rand.Intn(len(w.wordsFromTxt))
	return w.wordsFromTxt[r]
}
