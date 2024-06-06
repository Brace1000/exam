package main

import (
	"os"

	"github.com/01-edu/z01"
)
func main (){
if len(os.Args)== 0 {
	return
}
  args := os.Args
lastargs := args[len(os.Args)-1]
for _, ch := range lastargs{
	z01.PrintRune(ch)

}
z01.PrintRune('\n')
}