package main

import "fmt"

func validate(str string) bool {
	runes := []rune(str)
	count := 0
	for i, r := range runes {
		switch r {
		case '(':
			count++
		case ')':
			count--
			if count < 0 {
				fmt.Printf("Unexpected ')' at %d", i)
				return false
			}
		}
	}

	if count != 0 {
		fmt.Println("Not All Parens closed")
		return false
	} else {
		return true
	}
}
