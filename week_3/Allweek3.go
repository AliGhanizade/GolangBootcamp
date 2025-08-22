package week_3

import "fmt"

func SelectResult() {
	var practice int
	fmt.Print("Write Number of practice 1-Print 2-useChannel 0-All:")
	fmt.Scan(&practice)
	switch practice {
	case 1:
		runPrinterGoroutine()
	case 2:
		useChannel()
	}
}

func AllWeek3(){
	runPrinterGoroutine()
	useChannel()


}