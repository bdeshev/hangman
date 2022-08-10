package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var imputReader = bufio.NewReader(os.Stdin)

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
	hangmanState := 0
	for {
		printGameState(targetWord, guessedLetters, hangmanState)
		imput := readImput()
		if len(imput) != 1 {
			fmt.Println("Invalid Imput. Please use letters only...")
			continue
		}
	}
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
func printGameState(targetWord string, guessedLetters map[rune]bool, hangmanState int) {
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	fmt.Println()
	fmt.Println(getHangmanDrawing(hangmanState))
}

func getWordGuessingProgress(targetWord string, guessedLetters map[rune]bool) string {

	result := ""
	for _, ch := range targetWord {
		if ch == ' ' {
			result += " "
		} else if guessedLetters[unicode.ToLower(ch)] {
			result += fmt.Sprintf("%c", ch)
		} else {
			result += "_"
		}

		result += " "
	}
	return result
}

func getHangmanDrawing(hangmanState int) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("states/hangman%d", hangmanState))
	if err != nil {
		panic(err)
	}

	return string(data)
}

func readImput() string {
	fmt.Print("> ")
	imput, err := imputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(imput)
}
