package main

import (
	"fmt"
	"strconv"
)

// Function to take input of Dish w.r.t DishID
func OrderDish(d string) string {

	dishMap := map[string]string{
		"C": "Coffee",
		"D": "Dosa",
		"T": "Tomato soup",
		"I": "Idli",
		"V": "Vada",
		"B": "Bhature & Chhole",
		"P": "Paneer Pakoda",
		"M": "Manchurian",
		"H": "Hakka Noodle",
		"F": "French Fries",
		"J": "Jalebi",
		"L": "Lemonade",
		"S": "Spring Roll",
	}
	for key, value := range dishMap {
		if d == key {
			return value
		}

	}
	return "Invalid DishID. Please enter a valid one."
}

func OrderValue(d string) int {

	dishMap := map[string]int{
		"C": 40,
		"D": 80,
		"T": 20,
		"I": 20,
		"V": 25,
		"B": 30,
		"P": 30,
		"M": 90,
		"H": 70,
		"F": 30,
		"J": 30,
		"L": 15,
		"S": 20,
	}
	for key, value := range dishMap {
		if d == key {
			return value
		}

	}
	return 0
}

func main() {
	var order_quantity string
	var dishID string
	var temp int
	var earnings int

	fmt.Println("Hello Sir, How are you doing today?")
	fmt.Println("Welcome to XYZ Cafe")
	fmt.Println()
	fmt.Println("Our Menu")
	fmt.Println("--------")
	fmt.Println("1.  C : Coffee : Rs 40\n2.  D : Dosa : Rs 80\n3.  T : Tomato Soup : Rs 20\n4.  I : Idli : Rs 20\n5.  V : Vada : Rs 25\n6.  B : Bhature & Chhole : Rs 30\n7.  P : Paneer Pakoda : Rs 30\n8.  M : Manchurian : Rs 90\n9.  H : Hakka Noodle : Rs 70\n10. F : French Fries : Rs 30\n11. J : Jalebi : Rs 30\n12. L : Lemonade : Rs 15\n13. S : Spring Roll : Rs 20")

	for {
		fmt.Println()
		fmt.Println("Enter the dishID and order quantity :")

		fmt.Scan(&order_quantity)
		if order_quantity == "End" || order_quantity == "END" {
			break
		}
		fmt.Scan(&dishID)
		temp = OrderValue(dishID)

		if i, err := strconv.Atoi(order_quantity); err != nil {
			fmt.Println(err)
		} else {
			temp = temp * i
			fmt.Println(order_quantity, OrderDish(dishID), "is ordered and the total bill amount is", temp)
			fmt.Println("-----------------------------------------------------------------------------------")
			earnings = earnings + temp
		}
	}
	fmt.Println(" XYZ Cafe's total earning for the day is:", earnings)
}
