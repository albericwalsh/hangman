package mainHang

import (
	"fmt"
	"hangman/Game"
	"hangman/open"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args[1])
	var Save bool
	if !open.CheckPath(os.Args[1]) {
		fmt.Println("ERROR: " + os.Args[1] + " file not found")
	}
	if !open.CheckPath("./ressources/save.json") {
		fmt.Println("WARNING: no save file, creating one...")
		open.CreateSAveFile()
	} else {
		for _, Args := range os.Args[1:] {
			if strings.Contains(Args, "--startWith") {
				Save = true
			}
		}
		if Save {
			fmt.Println("You have a save file")
			Save := open.UnmarshalJson("./ressources/save.json")
			Game.Game(Save.Word, Save.Try, Save.Letters)
		} else {
			word := Game.RandomWord()

			w_exe := make([]rune, len(word))

			for i := 0; i < len(w_exe); i++ {
				w_exe[i] = '_'
			}

			Game.Game(word, 0, w_exe)
		}
	}
}
