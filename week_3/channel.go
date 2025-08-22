package week_3

import (
	"fmt"
	"sync"
)

func sendToChannel(chNumbers chan<- int, wg *sync.WaitGroup) {
	defer close(chNumbers)
	defer wg.Done()

	for i := 1; i < 11; i++ {
		chNumbers <- i
	}
	
}

func receiveChannelAndPrint(chNumbers <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for number := range chNumbers {
		fmt.Printf("%d ", number)
	}
}

func useChannel() {
	chNumbers := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go sendToChannel(chNumbers, &wg)
	go receiveChannelAndPrint(chNumbers, &wg)
	wg.Wait()
}
