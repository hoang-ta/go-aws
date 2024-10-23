package main

import (
	"fmt"
	test2 "go-aws/ticket1"
)

func main() {
	fmt.Println("hello world")
	t := test2.Ticket{
		ID: 1,
		Event: "test",
	}
	t.PrintEvent()
	fmt.Println(t)
	// t.PrintEvent()
}