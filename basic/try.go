package main

import "fmt"

func main() {
	/*=================== Map ======================*/
	datas := make(map[string]string) //map[key:string]value:string
	datas["firstname"] = "Tayathorn"
	datas["lastname"] = "Pan"
	datas["gender"] = "Female"

	// for k, value := range datas {
	// 	fmt.Printf("%v => %v\n", k, value)
	// }

	// //--------- Check data in map ----------
	fmt.Println("---------------")
	if datas["firstname"] == "" {
		fmt.Println(datas["firstname"])
	}

	if data, ok := datas["firstname"]; ok { // ok ข้างหลังคือ condition แบ่งโดย ; ข้างหน้าคือ statement
		fmt.Printf("\n%s : %t\n", data, ok)
	}

	/*============================================*/
}
