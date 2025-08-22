package week_3

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 28; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(5 * time.Millisecond)
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	for char := 'a'; char <= 'z'; char++ {
		fmt.Printf("%c ", char)
		time.Sleep(10 * time.Millisecond)
	}
}

func runPrinterGoroutine() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbers(&wg)
	go printLetters(&wg)
	wg.Wait()
}

