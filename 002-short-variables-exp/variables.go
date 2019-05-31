package main

import (
	"fmt"
)

func table() {
	x := 20
	for i := 1; i <= 20; i++ {
		fmt.Println(x, " X ", i, " = ", x*i)
	}
}

func main() {
	table()
}
