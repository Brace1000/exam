package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	first := os.Args[1]
	second := os.Args[2]
	secondIndex := 0
	for _, char := range first {
		found := false
		for i := 0; i <= len(second); i++ {
			if rune(second[i]) == char {
				found = true
				secondIndex++
				break
			}
		}
		if !found {
			return
		}

	}
	for _, char := range first {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
