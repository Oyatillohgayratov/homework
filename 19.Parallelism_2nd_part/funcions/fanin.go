package funcions

import (
    "sync"
)

func FanIn(intCh <-chan int, strCh <-chan string, floatCh <-chan float64) <-chan interface{} {
	out := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for val := range intCh {
			out <- val
		}
	}()

	go func() {
		defer wg.Done()
		for val := range strCh {
			out <- val
		}
	}()

	go func() {
		defer wg.Done()
		for val := range floatCh {
			out <- val
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
