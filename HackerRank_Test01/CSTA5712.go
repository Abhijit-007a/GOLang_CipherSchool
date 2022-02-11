// A new generation is created by all the cells in the board changing status depending on their neighboring cells. As neighboring cells
// considers all moving cells, both living and dead, are to be considered.

// The following describes the cell colony with '0' as living and 'X' as dead

// XXXXXX XX0XX0 0X0X0X 000000

// The cell's new status is determined by the following rules:

// • If the cell's current status is "alive":
// ○ At less than two living neighboring cells, the cell dies (subpopulation).
// ○ By two or three living neighboring cells, the cell will live on.
// ○ If the cell has more than three living neighboring cells, it will die (overpopulation)

// • If the cell is "dead":
// ○ The cell's status becomes "live" (reproduction) if it has exactly three living neighbors.
// Note that the update of the cell's status occurs at the same time! This means that we must determine the new status of all cells depending
// on the current status before the update actually occurs.

// Write a program which stimulates the cell life.

// Input Format

// 0

// Constraints

// nil

// Output Format

// nil

package main

import "fmt"

// func NextGeneration(univ [][]int, m, n int) [][]int {
// 	arr := [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
// 	for index, value := range arr {
// 		for index2, value2 := range value {
// 			fmt.Println(index, index2, value, value2)
// 		}
// 	}
// }

func main() {
	// 	fmt.Println("How many cycle do you want ot see ?")
	// 	var n int
	// 	fmt.Scanln(&n)
	// 	for i := 1; i <= n; i++ {
	// 		fmt.Println(NextGeneration(univ))
	// 	}
	// }
	arr := [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	m:=
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(arr[i][j])
		}
		fmt.Println(" ")
	}
}
