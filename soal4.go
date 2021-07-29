package main

import "fmt"

func main() {

	var inputs string
	fmt.Println("Input Kalimat : ")
	fmt.Scanln(&inputs)

	for i := len(inputs) - 1; i >= 0; i-- {
		fmt.Print(string(inputs[i]))
	}
}
