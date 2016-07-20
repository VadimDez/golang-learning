package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		var result string = "odd"

		if i % 2 == 0 {
			result = "even"
		}

		fmt.Println(i, result)
	}

	problem1()

	problem2()
}

func problem1()  {
	fmt.Println("Problem 1")

	for n := 3; n < 1000; n = n + 3 {
		fmt.Print(n, ",");
	}

	fmt.Println("")
}

func problem2() {

	fmt.Println("Problem 2")
	var isPrinted bool

	for c := 1; c < 100; c++ {
		isPrinted = false

		if c % 3 == 0 {
			fmt.Print("Fizz");
			isPrinted = true
		}

		if c % 5 == 0 {
			fmt.Print("Buzz");
			isPrinted = true
		}

		if !isPrinted {
			fmt.Print(c);
		}

		fmt.Print(",");
	}
	fmt.Println("")
}