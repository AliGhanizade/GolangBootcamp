package main

import (
	week2 "GolangBootcamp/week_2"
	week1 "GolangBootcamp/week_1"
	"fmt"
)

func main() {
	var Week int
	fmt.Print("Enter a Week(Number): ")
	fmt.Scanf("%d", &Week) // Reads an integer
	fmt.Println("You entered:", Week)
	switch Week {
	case 1:
		week1.AllWeek1()	
	case 2:
		week2.AllWeek2()
	}
}