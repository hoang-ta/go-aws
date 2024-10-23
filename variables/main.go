package main

import (
	"fmt"
)

func main() {
	var myName string = "Hoang"
	myName2 := "Hoang"
	myAge := 24

	fmt.Printf("Hello my name is %s and age is %d \n", myName, myAge)
	fmt.Printf("Hello my name 2 is %s \n", myName2)
	pets := [2]string{
		"bird",
		// "cat",
		"pandas",
	}

	pets[0] = "dog"
	// pets[1] = "cat"
	fmt.Println(pets)

	dynamic_pets := []string{"pandas", "wolf",}

	dynamic_pets2 := append(dynamic_pets, "cat")

	fmt.Println(dynamic_pets, dynamic_pets2)

	// dynamic_pets = slices.Delete(dynamic_pets, 0, len(dynamic_pets))
	// dynamic_pets = slices.Delete(dynamic_pets, 0, 10)
	fmt.Println(dynamic_pets)

}