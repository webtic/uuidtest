package main

import "github.com/nu7hatch/gouuid"

//import "fmt"

func main() {
	for i := 0; i < 10000000; i++ {
		_, _ = uuid.NewV4()

	}
}
