package main

import (
	"fmt"
)


type Stack struct {
	items []rune
}


func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}


func (s *Stack) Pop() (rune, bool) {
	if len(s.items) == 0 {
		return ' ', false 
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}


func isPalindrome(word string) bool {
	var stack Stack

	
	for _, char := range word {
		stack.Push(char)
	}

	
	var reversed string
	for i := 0; i < len(word); i++ {
		char, ok := stack.Pop()
		if ok {
			reversed += string(char)
		}
	}

	
	return word == reversed
}

func main() {
	
	words := []string{"radar", "arara", "golang", "level", "hello"}

	for _, word := range words {
		if isPalindrome(word) {
			fmt.Printf("\"%s\" é um palíndromo\n", word)
		} else {
			fmt.Printf("\"%s\" NÃO é um palíndromo\n", word)
		}
	}
}
