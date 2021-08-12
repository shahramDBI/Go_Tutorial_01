package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	a,b :=generateRandomNumbers()
	sum := add2(a, b)
	printNumber2(sum)
}

func generateRandomNumbers() (r1 int, r2 int)  {
	randomNumber1 := rand.Intn(10)
	randomNumber2 := rand.Intn(10)
	return randomNumber1, randomNumber2
}

func add2(num1 int, num2 int) (sum int)  {
	sum = num1 + num2
	return sum
}

func printNumber2(number int)  {
	fmt.Println("Sum is:", number)
	fmt.Printf("Sum is: %v", number)
}
