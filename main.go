package main

import (
	"fmt"
	"math/rand"
	"time"
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
	guessedLetters := map[rune]bool{}
	guessedLetters[rune(targetWord[0])] = true
	guessedLetters[rune(targetWord[len(targetWord)-1])] = true
	fmt.Println(targetWord)

	printGameState(targetWord, guessedLetters)
}

func getRandomWord() string {
	targetWord := dictionary[rand.Intn(len(dictionary))]
	return targetWord
}
func printGameState(targetWord string, guessedLetters map[rune]bool) {
	for _, ch := range targetWord {
		if ch == ' ' {
			fmt.Print(" ")
		} else if guessedLetters[ch] {
			fmt.Printf("%c", ch)
		} else {
			fmt.Printf("_")
		}
		fmt.Print(" ")
	}
	fmt.Println()
}
