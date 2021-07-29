package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var jam string
	var waktu string
	var output string
	var angka int

	// inputs = "04:01:00 PM"
	fmt.Print("Input Waktu : ")
	fmt.Scan(&waktu)
	fmt.Scan(&jam)

	pisah_waktu := strings.Split(waktu, ":")

	angka, err := strconv.Atoi(pisah_waktu[0])
	if err != nil {
		panic(err)
	}

	if jam == "PM" {
		if angka <= 12 {
			output += strconv.Itoa(angka + 12)
		} else {
			output += strconv.Itoa(angka)
		}
	} else if jam == "AM" {
		if angka == 12 {
			output += "00"
		} else {
			output += pisah_waktu[0]
		}
	}
	fmt.Println(output + ":" + pisah_waktu[1])
}
