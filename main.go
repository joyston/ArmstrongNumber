package main

import (
	"fmt"
	"strconv"
)

func cube(x int) int {
	return x * x * x
}

func IsArmstrongByString(a []int) {
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

func IsArmstrongWithoutString(a ...int) { //Variadic function
	for _, v := range a {
		sum := 0
		t := v
		for {
			if t == 0 {
				if v == sum {
					fmt.Printf("%v is an Armstrong\n", v)
				} else {
					fmt.Printf("%v is not an Armstrong\n", v)
				}
				break
			}
			s := t % 10
			sum += s * s * s
			t /= 10
		}
	}
}

func main() {
	var n, t int

	fmt.Println("Enter the number of test cases:")
	fmt.Scanln(&n)

	var a []int
	for i := 0; i < n; i++ {
		fmt.Println("Enter the number to test:")
		fmt.Scanln(&t)
		a = append(a, t)
	}

	// IsArmstrongByString(a)
	IsArmstrongWithoutString(a...)
}
