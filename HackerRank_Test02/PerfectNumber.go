package main

import "fmt"

func Perfect(n int) int {
	sum1 := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum1 += i
		}
	}
	if sum1 == n {
		return 1
	} else {
		return 0
	}
}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(Perfect(n))
}
