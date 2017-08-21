package main

import "fmt"

func main() {
	/*================== Array ==================*/

	// var datas [10]string // ตัวแปร data , length = 10 , type string
	// datas[0] = "Zero"
	// datas[1] = "One"
	// datas[2] = "Two"
	// datas[3] = "Three"
	// datas[4] = "Four"
	// datas[5] = "Five"
	// datas[6] = "Six"

	// for _, v := range datas {
	// 	fmt.Printf("%v\n", v)
	// }

	// fmt.Println(datas[1])
	// fmt.Println(datas[:])
	// fmt.Println(datas[2:5]) // print range of array

	//--------- Commenting in array ----------
	x := [4]int{
		1,
		2,
		3,
		// 4,
		5,
	}

	fmt.Println(len(x))
	fmt.Println(x[:])
	/*============================================*/
}
