package main

import "fmt"

const name = "Hello"

func main() {
	// variable type
	/*============================================*/
	// i := 1
	// f := 1.0
	// s := "Hello"
	// b := true

	// fmt.Printf("Type %T has value %d\n", i, i)
	// fmt.Printf("Type %T has value %f\n", f, f)
	// fmt.Printf("Type %T has value %s\n", s, s)
	// fmt.Printf("Type %T has value %t\n", b, b)
	/*============================================*/

	// fmt.Printf("%T => %v\n", name, name)

	/*============== Group Variable ==============*/
	i, j := 1, 2

	fmt.Printf("%T => %v\n", i, i)
	fmt.Printf("%T => %v\n", j, j)
	/*============================================*/
}
