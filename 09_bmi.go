package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// reader -> input
var reader = bufio.NewReader(os.Stdin)
// const
const mainTitle = "BMI Calculator!"
const separator = "---------------------------"
const weightPrompt = "Please enter your weight (kg): "
const heightPrompt = "Please enter your height (m): "
func main()  {
	// Output information
	fmt.Println(mainTitle)
	fmt.Println(separator)
	// Prompt for user input (weight + height)
	fmt.Print(weightPrompt)
	weightInput, _ := reader.ReadString('\n')
	fmt.Print(heightPrompt)
	heightInput, _ := reader.ReadString('\n')
	// Save that user input in variables
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)
	weight, _ := strconv.ParseFloat(weightInput, 64)
	heght, _ := strconv.ParseFloat(heightInput, 64)
	// Calculate the BMI (weight / (height * height)
	bmi := weight / (heght * heght)
	// Output the calculated BMI
	fmt.Printf("Your BMI: %.2f", bmi)
}
