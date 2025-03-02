package main

import "fmt"

func main() {
	// Initialize variables
	name := ""
	weight := 0.0
	height := 0.0

	// Prompt user for input
	fmt.Println("Enter your name:")
	fmt.Scanf("%s", &name)
	fmt.Println("Enter your weight in kg:")
	fmt.Scanf("%f", &weight)
	fmt.Println("Enter your height in meters:")
	fmt.Scanf("%f", &height)

	// Calculate BMI
	bmi := weight / (height * height)

	// Print the result
	fmt.Printf("Your BMI is: %.2f \n", bmi)
	status := ""
	if bmi < 18.5 {
		status = "Underweight"
	}
	if bmi >= 18.5 && bmi < 25 {
		status = "Normal weight"
	}
	if bmi >= 25 && bmi < 30 {
		status = "Overweight"
	}
	if bmi >= 30 {
		status = "Obese"
	}

	fmt.Println("Health Status:", status)

}
