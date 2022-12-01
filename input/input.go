package Input

import (
	"bufio"
	"fmt"
	"os"
)

// Input function
func Input() string {
	fmt.Print("Enter a character: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return GetInput(scanner.Text())
}

// Check the length of the input and return an index wich will be used to know what to do
func CheckLenInput(s string) int {
	if len(s) == 1 {
		if s >= "a" && s <= "z" {
			return 1
		} else if s >= "A" && s <= "Z" {
			return -1
		} else {
			fmt.Println("Enter an alpha character")
		}
	} else if len(s) > 1 {
		if s == "STOP" {
			return 0
		} else if s >= "a" && s <= "z" {
			return 1
		} else if s >= "A" && s <= "Z" {
			return -1
		} else {
			fmt.Println("Enter an alpha character")
		}
	}
	return 1
}

// Get the index and return the input
func GetInput(s string) string {
	word := []rune(s)
	switch CheckLenInput(s) {
	case 0:
		return "STOP" // stop and save the game
	case -1:
		return s // return the input correctly
	case 1:
		for index, i := range word {
			if i >= 97 && i <= 122 {
				word[index] = word[index] - 32
			} else if i >= 65 && i <= 90 {
				word[index] = word[index]
			}
		}
		return string(word)
	default:
		return GetInput(Input()) // retry
	}

}
