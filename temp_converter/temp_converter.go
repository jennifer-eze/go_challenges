package main

import "fmt"

func main() {
	// prompt user for temp converter type
	fmt.Print("type 1 to convert Celsius to Fahrenheit OR type 2 to convert Fahrenheit to Celsius: ")
	var convType float32
	fmt.Scanln(&convType)

	fmt.Print("Enter the temperature value you want converted: ")
	var tempValue float32
	var cel2Fah float32 = (tempValue - 32) * 5 / 9
	var fah2Cel float32 = (tempValue * 9 / 5) + 32
	fmt.Scanln(&tempValue)

	switch convType {
	case 1:
		fmt.Printf("%.2f converted to Fahrenheit is %.1f degrees", tempValue, cel2Fah)
	case 2:
		fmt.Printf("%.2f converted to Celsius is %.1f degrees", tempValue, fah2Cel)
	}

}
