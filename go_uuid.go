package main

import "code.google.com/p/go-uuid/uuid"
//import "fmt"

func main() {


	for i := 0; i < 10000000; i++ {
		//u := uuid.NewUUID()
		//fmt.Println(u)
		_ = uuid.NewUUID()

	}


}
