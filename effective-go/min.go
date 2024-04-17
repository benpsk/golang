package main

import "fmt"

func Min(a ...int) int {
	min := int(^uint(0) >> 1)

	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return 1
}

func init() {
	fmt.Println(Min(3, 5, 1, 7, 9))

	fmt.Println(Min())

	numbers := []int{8, 3, 6, 2, 1}
	fmt.Println(Min(numbers...))

}
