package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Initialize variables
	name := ""
	weight := 0.0
	height := 0.0

	// Prompt user for input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")
	name, _ = reader.ReadString('\n')
	fmt.Println("Enter your weight in kg:")
	fmt.Scanf("%f", &weight)
	fmt.Println("Enter your height in meters:")
	fmt.Scanf("%f", &height)

	fmt.Printf("\nWelcome, %s! Let's calculate your BMI.\n", name)
	fmt.Printf("Formula used: BMI = weight (kg) / (height (m) * height (m))\n\n")

	// Calculate BMI
	bmi := weight / (height * height)

	// Print the result
	fmt.Printf("Your BMI is: %.2f \n", bmi)
	status := ""
	if bmi < 18.5 {
		status = "Underweight"
	} else if bmi >= 18.5 && bmi < 25 {
		status = "Normal weight"
	} else if bmi >= 25 && bmi < 30 {
		status = "Overweight"
	} else if bmi >= 30 {
		status = "Obese"
	} else {
		status = "Invalid BMI"
	}

	fmt.Println("Health Status:", status)

}
