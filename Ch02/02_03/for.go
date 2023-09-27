// "for" loop examples
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("--1--")
	for i := 0; i < 3; i++ {
		if i < 1 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println()

	fmt.Println("--2--")
	for i := 0; i < 3; i++ {
		if i < 1 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("--3--")

	// go version of while loop
	a := 0
	for a < 3 {
		fmt.Println(a)
		a++
	}

// exlanation of the code below
// 	this is an infinite loop that will break when b > 2
// same a saying while true

	fmt.Println("--4--")
	b := 0
	for {
		if b > 2 {
			break
		}
		fmt.Println(b)
		b++
	}
}
