package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var number = 44
	for {
		var n = rand.Intn(100) + 1
		if n < number {
			fmt.Println("%v is too small.\n", n)
		} else if n > number {
			fmt.Println("%v is too big.\n", n)
		} else {
			fmt.Println("You got it! %v\n", n)
			break
		}
	}
}

