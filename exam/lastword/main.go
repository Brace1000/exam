package main

import (
	"os"

	"github.com/01-edu/z01"
)
func main () {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	lastword := ""

for i := len(args)-1; i>= 0;i--{
	if args[i]!= ' ' {
		lastword =string(args[i]) + lastword
	} else {
		if lastword != ""{
			break
		}
	}

}
for _, ch:= range lastword{
	z01.PrintRune(ch)
}
z01.PrintRune('\n')
}