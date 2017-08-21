package main

import "fmt"

func main() {
	/*=================== Map ======================*/
	datas := make(map[string]string) //map[key:string]value:string
	datas["firstname"] = "Tayathorn"
	datas["lastname"] = "Pan"
	datas["gender"] = "Female"

	for k, value := range datas {
		fmt.Printf("%v => %v\n", k, value)
	}
	/*============================================*/
}
