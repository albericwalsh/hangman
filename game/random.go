package Game

import (
	"hangman/open"
	"math/rand"
	"time"
)

func RandomWord() string {
	str := open.Open()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	lens := r1.Intn(len(str))
	return ConvertToMaj(str[lens])
	// return str[lens]
}

func ConvertToMaj(s string) string {
	str := ""
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			str += string(s[i] - 32)
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			str += string(s[i])
		}
	}
	return str
}
