package main

import "fmt"

func main() {

	// defining integer array
	var numbers []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Before Rotate:", numbers)

	// calling a function to rotate an array
	leftRotate(numbers, 3)

	fmt.Println("After Rotate:", numbers)
}

// function rotate based on the count
func leftRotate(arr []int, count int) {
	for i := 0; i < count; i++ {
		leftRotatebyOne(arr) // calling function to perform rotation ont time
	}
}

// function perform left rotate only one time
func leftRotatebyOne(arr []int) {

	var i int = 0

	var temp int = arr[0]

	for ; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}

	arr[i] = temp
}
