package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const gr = 100
	var m sync.Mutex
	var wg sync.WaitGroup
	wg.Add(gr * 2)

	var n int = 0
	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			n++
			m.Unlock()
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			n--
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final value of n:", n)
	// Check race condition by running with -race flag
}
