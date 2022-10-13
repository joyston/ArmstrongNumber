package main

import (
	"fmt"
	"strconv"
)

func cube(x int) int {
	return x * x * x
}

func main() {
	var n, t int

	fmt.Println("Enter the number of test cases:")
	fmt.Scanln(&n)

	var a []int
	for i := 0; i < n; i++ {
		fmt.Scanln(&t)
		a = append(a, t)
	}

	var s string

	for _, v := range a {
		s = fmt.Sprint(v)
		sum := 0
		for j := 0; j < len(s); j++ {
			x, _ := strconv.Atoi(string(s[j]))
			sum = sum + cube(x)
		}
		if sum == v {
			fmt.Println("Armstrong number")
		} else {
			fmt.Println("Not an Armstrong number")
		}
	}
}
