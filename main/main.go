package main

import (
	// "fmt"
	// "hangman/input"
	"hangman/ascii"
	// "hangman/open"
	// "hangman/parth"
	// "hangman/RGB"
)

func main() {
	// Getting the input from the user and storing it in the variable s.
	// s := input.GetInput(input.Input())
	Line := [][]string{ascii.ConvertToAsciiArt('_'),ascii.ConvertToAsciiArt('O'),ascii.ConvertToAsciiArt('Z'),ascii.ConvertToAsciiArt('U')}
	ascii.PrintLine(Line)
}
