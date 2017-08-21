package main

import "fmt"

func main() {
	/*============== Slice =======================*/

	// var datas = make([]string, 3)
	// datas[0] = "Zero"
	// datas[1] = "One"
	// datas[2] = "Two"

	// fmt.Println(datas[:])

	//--------- Slice : Append, Copy ----------

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Printf("%v\n", slice1)
	fmt.Printf("%v\n", slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
	/*============================================*/
}
