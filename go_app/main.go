// Made by Mohamad
// Made on April 6th 2023
//
// This program calculates the volume of a sphere

package main

import (
	"fmt"
	"math"
)

func main() {
	// This function calculates the volume of a sphere
	var radius float64
	var volume float64

	// Input
	fmt.Print("Enter the radius (cm): ")
	fmt.Scan(&radius)

	// Process
	volume = (4.0 / 3.0) * math.Pi * math.Pow(radius, 3)

	// Round area
	roundedVolume := fmt.Sprintf("%.2f", volume)

	// Output the volume to the user
	fmt.Println("\nThe volume is", roundedVolume, "cm³")

	// Round the volume to 2 decimal places
	fmt.Println("\nDone.")
}
