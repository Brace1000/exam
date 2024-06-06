package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// if len(os.Args) < 1 {
	// 	return
	// }
	 count := 0
	for i := 0; i < len(os.Args[1:]); i++ {
		count++
	}
	str := ""
	for count > 0 {
		digit := count % 10
		str = string(rune(digit+'0')) + str
		count /= 10
	}

	for _, val := range str {
		z01.PrintRune(val)
	}
	z01.PrintRune('\n')
}