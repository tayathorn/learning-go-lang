package main

import "fmt"

func main() {
	/*============== Slice =======================*/

	var datas = make([]string, 3)
	datas[0] = "Zero"
	datas[1] = "One"
	datas[2] = "Two"

	fmt.Println(datas[:])
	/*============================================*/
}
