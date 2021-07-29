package main

import (
	"fmt"
	"strings"
)

func main() {

	var inputs string
	var output string

	fmt.Println("Input Kalimat : ")
	fmt.Scanln(&inputs)

	for i := len(inputs) - 1; i >= 0; i-- {
		output += string(inputs[i])
	}

	if strings.ToUpper(inputs) == strings.ToUpper(output) {
		fmt.Println("Result True")
	} else {
		fmt.Println("Result False")
	}

}
