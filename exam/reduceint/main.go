package main

import (
	"fmt"
)

func ReduceInt(a []int, f func(int, int) int) {
	acc := a[0]
	for _, val := range a {
		acc = f(acc, val)
	}
	fmt.Println(acc)
	fmt.Println()
}
