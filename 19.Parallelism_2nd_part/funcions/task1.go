package funcions

import "sync"

func DoTask1(out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- 1
}
