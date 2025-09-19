package week_6

import (
	"fmt"
	"sync"
)


func SelectResult() {
	var practice int
	fmt.Print("Write Number of practice :\n\t1-web crawler \n\t2-web socket echo \n\t3-web socket chat \n")
	fmt.Scan(&practice)
	switch practice {

	case 1:
		var wg sync.WaitGroup
		wg.Add(2)
		go getWebLink(&wg)
		go getWebTitle(&wg)
		wg.Wait()
	case 2:
		StartWebSocketEchoServer()
	case 3 : 
		StartChatServer()
	}
}