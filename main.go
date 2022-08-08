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
	for _, ch := range targetWord {
		if ch == ' ' {
			fmt.Print(" ")
		} else if guessedLetters[unicode.ToLower(ch)] {
			fmt.Printf("%c", ch)
		} else {
			fmt.Printf("_")
		}
		fmt.Print(" ")
	}
	fmt.Println()
	fmt.Println()
	fmt.Println(getHangmanDrawing(1))
}

func getHangmanDrawing(hangmanState int) string {
	data, err := ioutil.ReadFile("states/hangman1")
	if err != nil {
		panic(err)
	}

	return string(data)
}
