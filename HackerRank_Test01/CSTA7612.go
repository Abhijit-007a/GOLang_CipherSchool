// Find the difference between the square of the sum and the sum of the squares of the first N natural numbers.

// Let the input be 10 The square of the sum of the first 10 natural numbers is (1 + 2 + ... + 10)² = 55² = 3025.

// The sum of the squares of the first 10 natural numbers is 1² + 2² + ... + 10² = 385.

// Hence the difference between the square of the sum of the first ten natural numbers and the sum of the squares of the first ten natural numbers is 3025 - 385 = 2640.

// Functions naming format:-

// func SquareOfSum(n int) int {} func SumOfSquares(n int) int {} func Difference(n int) int {}

// Input Format

// 10

// Constraints

// 0

// Output Format

// 2640

package main

import "fmt"

func SumOfSquares(n int) int {
	sumsquare := 0
	for i := 1; i < n+1; i++ {
		sumsquare = sumsquare + (i * i)
	}
	return sumsquare
}

func SquareOfSum(n int) int {
	squaresum := 0
	for i := 1; i < n+1; i++ {
		squaresum = squaresum + i
	}
	return squaresum * squaresum
}

func Difference(n int) int {

	Diff := SquareOfSum(n) - SumOfSquares(n)
	return Diff
}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(Difference(n))
}
