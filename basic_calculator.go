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
	fmt.Println("Choose your operation: addition, subtraction, multiplication or division")
	var operationType string
	var value1Float, err = strconv.ParseFloat(value1, 64)
	var value2Float, err2 = strconv.ParseFloat(value2, 64)
	fmt.Scanln(&operationType)
	if err == nil && err2 == nil {
		if operationType == "addition" {
			var addition float64 = value1Float + value2Float
			fmt.Printf("The sum of %.2f and %.2f is %.2f", value1Float, value2Float, addition)
		} else if operationType == "subtraction" {
			var subtraction float64 = value1Float - value2Float
			fmt.Printf("The sum of %.2f and %.2f is %.2f", value1Float, value2Float, subtraction)
		} else if operationType == "multiplication" {
			var multiplication float64 = value1Float * value2Float
			fmt.Printf("The sum of %.2f and %.2f is %.2f", value1Float, value2Float, multiplication)
		} else if operationType == "division" {
			var division float64 = value1Float / value2Float
			fmt.Printf("The sum of %.2f and %.2f is %.2f", value1Float, value2Float, division)
		}
	} else {
		fmt.Println("Error:", err, err2)
	}
}
