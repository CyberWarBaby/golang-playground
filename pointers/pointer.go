package main

import "fmt"

func main() {
	age := 32

	agePointer := &age

	fmt.Println("Age: ", *agePointer)

	editAdultYears(agePointer)
	fmt.Println(age)
}

func editAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
