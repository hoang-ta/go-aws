package main

import (
	"fmt"
)

func main() {
	dynamic_pets := []string{"pandas", "wolf",}


	for i:=0; i< len(dynamic_pets); i++ {
		fmt.Printf("Animal %s\n", dynamic_pets[i])
	}

	for index, value := range dynamic_pets {
		fmt.Printf("range %d , %s\n", index, value)
	}

	for value := range 10 {
		fmt.Println(value)
	}

	i := 0

	for i < 5 {
		fmt.Println("i: ",i)
		i++
	}
}