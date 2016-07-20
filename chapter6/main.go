package main

import "fmt"

func main() {
	x := [5]float64{ 98, 93, 77, 82, 83 }

	var total float64 = 0
	var arrayLength = len(x)
	for i := 0; i < arrayLength; i++ {
		total += x[i]
	}
	fmt.Println(total / float64(arrayLength))

	total = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))

	// slices

	y := make([]float64, 5, 10)
	fmt.Println(y)
	z := x[0:5]
	fmt.Println(z)


	slice1 := []int{ 1, 2, 3 }
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice2)

	slice3 := []int{1,2,3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)

	// Maps

	m := make(map[string]int)
	m["key"] = 10
	delete(m, "key")
	fmt.Println(m)


	//
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	fmt.Println(elements)

	arr := [6]string{"a","b","c","d","e","f"}
	fmt.Println(arr[2:5])

	fmt.Println("max value is: ", findMax([]int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}))
}

func findMax(x []int) int {
	var max int = x[0]

	for _, value := range x {
		if value > max {
			max = value
		}
	}

	return max
}