package main

import "fmt"

const str string = "Hello World"

func main1() {
	fmt.Println(str)

	x := "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)

	x = "first "
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)

	var (
		a = "hello"
		b = "world"
	)
	fmt.Println(a == b)
}

func main2() {
	fmt.Println("Enter a number: ")

	var input float64

	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}

func fahrenheit() {
	fmt.Println("Enter Fahrenheit temperature: ")

	var f float64

	fmt.Scanf("%f", &f)

	fmt.Println((f - 32) * 5/9, " Celsius")
}

func main() {
	const ration float64 = 0.3048

	fmt.Println("Enter feets amount: ")

	var feets float64

	fmt.Scanf("%f", &feets)

	fmt.Println(feets, " feets = ", feets * ration, " meters")
}