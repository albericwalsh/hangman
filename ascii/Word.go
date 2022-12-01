package ascii

import (
	"fmt"
	"hangman/open"
)

//open standar.txt and parse the ascii art file

func ConvertToAsciiArt(r rune) []string {
	var Char [][]string
	index := 0
	if r >= 65 && r <= 90 {
		index = int(r) - 65
	} else if r == '_' {
		index = 26
	}
	j := 0
	str := open.OpenPath("./ascii/standard.txt")
	Char = append(Char, []string{})
	for i := range str {
		if (i+1)%9 != 0 {
			Char[j] = append(Char[j], str[i])
		} else {
			Char[j] = append(Char[j], str[i])
			j++
			Char = append(Char, []string{})
		}
	}
	return Char[index]
}

func PrintLine(tab [][]string) {
	for i := 0; i < 8; i++ {
		for j := range tab {
			fmt.Print(tab[j][i])
		}
		fmt.Println()
	}
}


