package week_6

import (
	"fmt"
	"sync"
)


func SelectResult() {
	var practice int
	println("Write Number of practice 1-web craweler 0-All:")
	fmt.Scan(&practice)
	switch practice {

	case 1:
		var wg sync.WaitGroup
		wg.Add(2)
		go getWebLink(&wg)
		go getWebTitle(&wg)
		wg.Wait()
	case 0:
		
	}
}