package main

import "fmt"

func main() {
	// variable type
	i := 1
	f := 1.0
	s := "Hello"
	b := true

	fmt.Printf("Type %T has value %d\n", i, i) //integer
	fmt.Printf("Type %T has value %f\n", f, f) //float
	fmt.Printf("Type %T has value %s\n", s, s) //string
	fmt.Printf("Type %T has value %t\n", b, b) //boolean
}
