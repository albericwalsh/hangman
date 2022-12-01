package ascii

import (
	"fmt"
)

func HangmanPosition0() {
	fmt.Println(`
         
         
         
         
         
         
         
`)
}

func HangmanPosition1() {
	fmt.Println("Not present in the word, 9 attempts remaining")
	fmt.Println(`
         
         
         
         
         
         
=========
`)
}

func HangmanPosition2() {
	fmt.Println("Not present in the word, 8 attempts remaining")
	fmt.Println(`
         
      |  
      |  
      |  
      |  
      |  
=========
`)
}

func HangmanPosition3() {
	fmt.Println("Not present in the word, 7 attempts remaining")
	fmt.Println(`
  +---+  
      |  
      |  
      |  
      |  
      |  
=========
`)
}

func HangmanPosition4() {
	fmt.Println("Not present in the word, 6 attempts remaining")
	fmt.Println(`

  +---+  
  |   |  
      |  
      |  
      |  
      |  
=========
`)
}

func HangmanPosition5() {
	fmt.Println("Not present in the word, 5 attempts remaining")
	fmt.Println(`
  +---+  
  |   |  
  O   |  
      |  
      |  
      |  
=========
`)
}

func HangmanPosition6() {
	fmt.Println("Not present in the word, 4 attempts remaining")
	fmt.Println(`
  +---+  
  |   |  
  O   |  
  |   |  
      |  
      |  
=========
`)
}

func HangmanPosition7() {
	fmt.Println("Not present in the word, 3 attempts remaining")
	fmt.Println(`

  +---+  
  |   |  
  O   |  
 /|   |  
      |  
      |  
=========
`)
}

func HangmanPosition8() {
	fmt.Println("Not present in the word, 2 attempts remaining")
	fmt.Println(`
  +---+  
  |   |  
  O   |  
 /|\  |  
      |  
      |  
=========
`)
}

func HangmanPosition9() {
	fmt.Println("Not present in the word, 1 attempts remaining")
	fmt.Println(`
  +---+  
  |   |  
  O   |  
 /|\  |  
 /    |  
      |  
=========
`)
}

func HangmanPosition10() {
	fmt.Println("Not present in the word, josé is dead")
	fmt.Println(`

  +---+  
  |   |  
  O   |  
 /|\  |  
 / \  |  
      |  
=========
`)
}
