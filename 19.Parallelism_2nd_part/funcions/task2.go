package funcions

import (
    "sync"
)

func DoTask2(out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- "hello"
}