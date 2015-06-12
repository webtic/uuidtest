package main

import "github.com/twinj/uuid"

import "fmt"

func main() {

	
	//fmt.Println(u)
	for i := 0; i < 10000000; i++ {
		u:= uuid.NewV4()
		fmt.Println(u)

	}
	//fmt.Println(u)

}
