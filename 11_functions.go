package main

import (
	"fmt"
)

func main()  {
	a := 10
	b := 5
	sum := add1(a, b)
	printNumber1(sum)
}

func add1(num1 int, num2 int) int  {
	return num1 + num2
	//return sum
}

func printNumber1(number int)  {
	fmt.Println("Sum is:", number)
	fmt.Printf("Sum is: %v", number)
}
