// Created by: Ryan-Shaw-2
// Created on: May 2021
//
// This program calculates the area of a trapezoid

package main

import "fmt"

func main() {
	// This function calculates the area of a trapezoid
	var aBase int
	var bBase int
	var height int

	// input
	fmt.Println("This program calculates the area of a trapezoid")
	fmt.Println()
	fmt.Print("Enter the length of a base (cm): ")
	fmt.Scanln(&aBase)
	fmt.Print("Enter the length of b base (cm): ")
	fmt.Scanln(&bBase)
	fmt.Print("Enter the height (cm): ")
	fmt.Scanln(&height)
	fmt.Println()

	// process
	var area = ((aBase + bBase) / 2) * height

	// output
	fmt.Println("The area is:", area, "cm²")
}
