package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
	"unicode"
)

var dictionary = []string{
	"Zombie",
	"Gopher",
	"United States of America",
	"Indonesia",
	"Nazism",
	"Apple",
	"Programming",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	targetWord := getRandomWord()
	guessedLetters := initializeGuessedWord(targetWord)

	printGameState(targetWord, guessedLetters)
}

func initializeGuessedWord(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true

	return guessedLetters
}

func getRandomWord() string {
	targetWord := dictionary[rand.Intn(len(dictionary))]
	return targetWord
}
func printGameState(targetWord string, guessedLetters map[rune]bool) {
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	fmt.Println()
	fmt.Println(getHangmanDrawing(1))
}

func getWordGuessingProgress(
	targetWord string,
	guessedLetters,
	map[rune]bool)
	string
 {
	result := ""
	for _, ch := range targetWord {
		if ch == ' ' {
			result += " "
			} else if guessedLetters[unicode.ToLower(ch)] {
			result += fmt.SPrintf("%c", ch)
		} else {
			result += " "
		}
		result += " "
}

func getHangmanDrawing(hangmanState int) string {
	data, err := ioutil.ReadFile("states/hangman1")
	if err != nil {
		panic(err)
	}

	return string(data)
}
