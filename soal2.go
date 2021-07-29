package main

import (
	"fmt"
	"strings"
)

func main() {

	var inputs string
	// var output string

	// inputs = "halosssssss@.id"
	fmt.Print("Input email : ")
	fmt.Scanln(&inputs)

	if (strings.Contains(inputs, ".co.id") != true) && (strings.Contains(inputs, ".id") != true) {
		fmt.Println("Bukan Email")
	} else {
		if strings.Contains(inputs, "@") != true {
			fmt.Println("Bukan Email")
		} else {
			hitung_panjang := strings.Split(inputs, "@")
			if len(hitung_panjang[0]) > 20 {
				fmt.Println("Bukan Email")
			} else {
				if strings.Contains(hitung_panjang[1], ".") != true {
					fmt.Println("Bukan Email")
				} else {
					fmt.Println("Ini adalah Email")
				}
			}
		}
	}
}
