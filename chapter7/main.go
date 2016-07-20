package main

import (
	"fmt"
)

func main() {
	x, y := f()

	fmt.Println(x, y)

	fmt.Println(add(10, 30, 20))

	closure := func(a, b int) int {
		return a + b
	}

	fmt.Println(closure(1, 1))

	//
	fmt.Println("Even")

	nextEven := makeEvenGenerator()

	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	//
	fmt.Println("factorial of 4 is", factorial(4))


	// Defer
	//f, _ := os.Open(filename)
	//defer f.Close()

	// Panic
	//defer func() {
	//	str := recover()
	//	fmt.Println(str)
	//}()
	//panic("PANIC")


	// problem
	fmt.Println("fibonacci value of 3:", fibonacci(3))
	fmt.Println("fibonacci value of 7:", fibonacci(7))
}

func f() (int, int) {
	return 5, 6
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(i uint) uint {
	if i == 0 {
		return 1
	}

	return i * factorial(i - 1)
}

func fibonacci(i uint) uint {
	if i < 0 {
		return 0
	}

	if i <= 1 {
		return i
	}

	return fibonacci(i - 1) + fibonacci(i - 2)
}