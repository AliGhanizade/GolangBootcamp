package main

import (
	week1 "GolangBootcamp/week_1"
	week2 "GolangBootcamp/week_2"
	week3 "GolangBootcamp/week_3"
	week4 "GolangBootcamp/week_4"
	week5 "GolangBootcamp/week_5"
	week6 "GolangBootcamp/week_6"
	"fmt"
)

// run main and select number of week and you can see result
func main() {
	var Week int
	fmt.Print("Enter a Week(Number) : ")
	fmt.Scanf("%d", &Week) // Reads an integer
	fmt.Println("You entered:", Week)
	switch Week {
	case 1:
		week1.SelectResult()
	case 2:
		week2.SelectResult()
	case 3:
		week3.SelectResult()
	case 4:
		week4.Run()
	case 5:
		week5.SelectResult()
	case 6:
		week6.SelectResult()
	default:
		fmt.Println("Invalid week number. Please enter a number.")
	}
}
