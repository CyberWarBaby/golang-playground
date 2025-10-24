package main

import "fmt"

type customString string

func (str customString) log() {
	fmt.Println(str)
}
func main() {
	var name customString
	name = "max"

	name.log()
}
