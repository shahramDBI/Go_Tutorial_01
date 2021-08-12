package main

import "fmt"

func main()  {
	age := 41
	fmt.Println(age)
	myAge := &age  // var myAge  *int  --> Address my value not Age *myAge = 41
	fmt.Println(myAge)
	fmt.Println(*myAge)
	doubledAge := double(age)
	fmt.Println(doubledAge)
}
func double(number int) int {
	result := number * 2
	number = 100
	return result
}
