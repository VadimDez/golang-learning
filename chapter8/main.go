package main

import (
	"fmt"
)

func main() {
	//x := 5
	//zero(&x)
	//fmt.Println(x);

	xPtr := new(int)
	zero(xPtr)
	fmt.Println(*xPtr);

	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)
}

func zero(xPtr *int) {
	*xPtr = 0;
}

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}