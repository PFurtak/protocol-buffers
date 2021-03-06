package main

import (
	"fmt"

	simplepb "github.com/basics/go/src/simple"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "Simple Test",
		SampleList: []int32{1, 2, 4, 8},
	}

	fmt.Println(sm)

	sm.Name = "I renamed you"
	fmt.Println(sm)

	fmt.Println("The ID is: ", sm.Id)
}
