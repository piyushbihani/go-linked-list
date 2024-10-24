// Go program to illustrate the
// use of logical operators
package main

import (
	"fmt"
)

func variadic_functions() {
	arr := []int{5, 10, 15, 20, 25,}
	fmt.Printf("Sum of numbers is: %d\n", add(arr...))
	fmt.Printf("Sum of numbers is: %d\n", add_(10, 20, 30, 40))
	fmt.Printf("Sum of numbers is: %d\n", add_(10))
	fmt.Printf("Sum of numbers is: %d\n", add_())
}
