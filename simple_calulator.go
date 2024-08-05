package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("This is a simple calculator app")
	fmt.Println("Enter first value:")
	var value1 string
	fmt.Scanln(&value1)
	fmt.Println("Enter second value:")
	var value2 string
	fmt.Scanln(&value2)
	fmt.Println("This calculator only performs addition.")
	var value1Float, err = strconv.ParseFloat(value1, 64)
	var value2Float, err2 = strconv.ParseFloat(value2, 64)
	var sum float64 = value1Float + value2Float
	if err != nil && err2 != nil {
		fmt.Println("Error:", err, err2)
	} else {
		fmt.Printf("The sum of %.2f and %.2f is %.2f\n", value1Float, value2Float, sum)
	}
}
