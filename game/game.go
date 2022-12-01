package Game

import (
	"fmt"
	"hangman/Input"
	"hangman/ascii"
	"hangman/open"
	"os"
)

func len_world(w_rep []rune, l rune) int {
	n := 0

	for _, c := range w_rep {
		if c == l {
			n++
		}
	}
	return n
}

func is_it(w_rep, w_exe []rune, l rune) int {
	for i := 0; i < len(w_rep); i++ {
		if w_exe[i] == '_' {
			if w_rep[i] == l {
				return i
			}
		}
	}
	return -1
}

func draw_hang(try int) {
	switch try {
	case 0:
		ascii.HangmanPosition0()
	case 1:
		ascii.HangmanPosition1()
	case 2:
		ascii.HangmanPosition2()
	case 3:
		ascii.HangmanPosition3()
	case 4:
		ascii.HangmanPosition4()
	case 5:
		ascii.HangmanPosition5()
	case 6:
		ascii.HangmanPosition6()
	case 7:
		ascii.HangmanPosition7()
	case 8:
		ascii.HangmanPosition8()
	case 9:
		ascii.HangmanPosition9()
	case 10:
		ascii.HangmanPosition10()
	}
}

func Game(word string, try int, letters []rune) {
	win := 0
	n := 0
	w_rep := []rune(word)
	w_exe := letters
	po := 0

	w_exe[len(word)-1] = w_rep[len(word)-1]
	fmt.Println("Good Luck, you have 10 attempts.")
	for win != 1 && win != -1 {
		Line := [][]string{}
		if try >= 10 {
			win = -1
		}
		if win == -1 {
			fmt.Printf("\x1bc")
			ascii.HangmanPosition10()
			os.Exit(0)
		}
		for i := 0; i < len(word); i++ {
			Line = append(Line, ascii.ConvertToAsciiArt(w_exe[i]))
		}
		fmt.Println()
		draw_hang(try)
		ascii.PrintLine(Line)
		l := Input.Input()
		if len(l) > 1 {
			if l == "STOP" {
				fmt.Println("You have stopped the game")
				Save := open.Save{Word: word, Try: try, Letters: w_exe}
				open.MarshalJson("./ressources/save.json", open.Save(Save))
				os.Exit(0)
			}
			rep := []rune(word)
			if l == string(rep) {
				win = 1
			} else {
				try += 2
			}
		} else {
			switch is_it(w_rep, w_exe, rune(l[0])) {
			case -1:
				fmt.Printf("\x1bc")
				try += 1
			default:
				fmt.Printf("\x1bc")
				po = len_world(w_rep, rune(l[0]))
				for m := 0; m < po; m++ {
					w_exe[is_it(w_rep, w_exe, rune(l[0]))] = rune(l[0])
				}
			}
			for k := 0; k < len(w_exe); k++ {
				if w_exe[k] == '_' {
					n++
				}
			}
			if n == 0 {
				win = 1
			} else {
				n = 0
			}
		}
	}
	if win == 1 {
		Line := [][]string{}
		for i := 0; i < len(w_exe); i++ {
			Line = append(Line, ascii.ConvertToAsciiArt(w_rep[i]))
		}
		ascii.PrintLine(Line)
		// fmt.Println(string(w_exe))
		fmt.Println("\nCongrats !")
		// remove save file
		os.Remove("./ressources/save.json")
	}
}
