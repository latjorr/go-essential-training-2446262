package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main() {
	getinput() // Accept user input
	dataTypes() // Print out the data types
	castingInGo() // Show how to cast types

	
}



func castingInGo() {
	// casting in go is done by using the type in parenthesis
	pl("---Casting----")
	var i int = 42
	pl(i, reflect.TypeOf(i))

	var f float64 = float64(i)
	pl(f, reflect.TypeOf(f))

	var u uint = uint(f)
	pl(u, reflect.TypeOf(u))
	
	myStr := "54042"
	myInt, err := strconv.Atoi(myStr)
	if err != nil {
		log.Fatal(err)
	}
	pl(myInt, err)
}

func dataTypes() {
	pl(reflect.TypeOf(42)) // int
	pl(reflect.TypeOf(3.1415)) // float64 (default for floating point numbers)
	pl(reflect.TypeOf(true)) // bool (true or false)
	pl(reflect.TypeOf("Hello, World!")) // string
	pl(reflect.TypeOf('a')) // rune (alias for int32) are used for unicode characters
	pl(reflect.TypeOf(1 + 2i)) // complex128 can store complex numbers with float64 precision
	pl(reflect.TypeOf([]int{1, 2, 3})) // []int. this is a slice
	pl(reflect.TypeOf('👍')) //rune (alias for int32) are used for unicode characters
}

func getinput() {
	// this is for taking input from the user
	pl("what is your name?")

	/* the  reader is a buffered reader, which means it will
	 * read the input in chunks, rather than one character at a time
	 * this is more efficient
	 */
	reader := bufio.NewReader(os.Stdin) 

	// read the input until the user hits enter (\n)
	name, err := reader.ReadString('\n')
	
	/* if there is an error, print it out and exit the program with a fatal error
	 * otherwise, print out the greeting
	*/
	if err != nil {
		log.Fatal(err)
	} else {
		pl("hello", name)
	}
}

