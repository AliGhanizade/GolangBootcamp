package main

import (
	week1 "GolangBootcamp/week_1"
	week2 "GolangBootcamp/week_2"
	week3 "GolangBootcamp/week_3"
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
	case 0:
		
	}
}
