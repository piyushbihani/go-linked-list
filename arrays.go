package main

import (
	"fmt"
	"bytes"
)

func array() {
	/*
		Arrays can be decalred using this syntax and we can initialize them later
		var arr [10]int
		arr[0] = 1
		We can also do shorthand decalaration
		arr2 := [10]int
	*/
	/*
	This is the way to declare and initialize a slice
	var arr = []int{10, 20}

	
	arr2 := make([]int, len(arr), 3)
	copy(arr2, arr)
	// arr2 = append(arr2, 40)
	fmt.Println(arr, cap(arr))
	fmt.Println(arr2, cap(arr2))
	*/

	/*
	arr := [3]int{1, 2, 3}
    fmt.Println("Array:", arr)

    // Slice from the array
    s := arr[:]
    fmt.Println("Slice before modification:", s, len(s))

    // Modify the slice
    arr[0] = 100
    fmt.Println("Array after modification:", arr)

    // The original array is also modified because the slice references it
    fmt.Println("Slice after array modification:", s)

    // Using append on a slice
    s = append(s, 4, 5)
    fmt.Println("Slice after append:", s, len(s))
    fmt.Println("Array after append:", arr)
	*/
	slice := []byte("        Hello World*********")
 
    fmt.Printf("Slice: %s", bytes.Trim(slice, "* d"))
}

// Go program to illustrate 
// how to compare two arrays 

// func array() { 

// // Arrays	 
// arr1:= [3]int{9,7,6} 
// arr2:= [...]int{9,7,6} 
// arr3:= [3]int{9,5,3} 

// // Comparing arrays using == operator 
// fmt.Println(arr1==arr2) 
// fmt.Println(arr2==arr3) 
// fmt.Println(arr1==arr3)

// arr4:= [4]int{9,7} 
// fmt.Println(arr1==arr4) 

// // This will give and error because the 
// // type of arr1 and arr4 is a mismatch 
// /* 
// arr4:= [4]int{9,7,6} 
// fmt.Println(arr1==arr4) 
// */
// } 
