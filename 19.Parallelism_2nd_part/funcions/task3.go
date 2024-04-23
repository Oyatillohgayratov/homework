package funcions

import (
    "sync"
)

func DoTask3(out chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- 3.14
}