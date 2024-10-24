package main

import (
	"fmt"
)

func add_(elements ...int)int{
	fmt.Print("")
	sum := 0
	for _, num := range elements{
		sum += num
	}
	return sum
}

func add(ele ...int)int{
	sum := 0
	for _, num := range ele{
		sum += num
	}
	return sum
}

