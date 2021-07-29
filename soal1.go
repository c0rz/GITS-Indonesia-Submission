package main

import (
	"fmt"
)

func main() {

	var inputs int
	fmt.Print("Input bilangan : ")
	fmt.Scanln(&inputs)

	if (inputs%3 == 0) && (inputs%5 == 0) {
		fmt.Println("Output : Hello World")
	} else if inputs%3 == 0 {
		fmt.Println("Output : Hello")
	} else if inputs%5 == 0 {
		fmt.Println("Output : World")
	}

}
