package main

import (
	"fmt"
	"strconv"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return fmt.Sprintf("%s\nYou have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.",
		Welcome(name), table, direction, distance, neighbor)
}

func main() {
	fmt.Println(Welcome("Christiane"))
	fmt.Println(HappyBirthday("Frank", 58))
	fmt.Println(AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))
}
