// 1. Take as input N, a number. Print odd numbers in decreasing sequence (up until 0) and even numbers in increasing sequence
// (up until N). E.g. for N = 6 print 5, 3, 1, 2, 4

package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	OddNumber(n)
	EvenNumber(n)
}

func EvenNumber(n int) {
	for i := 1; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(i)
		}
	}
}

func OddNumber(n int) {
	for i := n - 1; i > 0; i-- {
		if i%2 != 0 {
			fmt.Print(i)
		}
	}
}
