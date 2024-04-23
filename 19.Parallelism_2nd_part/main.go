package main

import (
	"fmt"
	"sync"
	f "homework/funcions"
)
func main() {
	intCh := make(chan int)
	strCh := make(chan string)
	floatCh := make(chan float64)

	var wg sync.WaitGroup
	wg.Add(3)

	go f.DoTask1(intCh, &wg)
	go f.DoTask2(strCh, &wg)
	go f.DoTask3(floatCh, &wg)

	results := f.FanIn(intCh, strCh, floatCh)

	go func() {
		wg.Wait()
		close(intCh)
		close(strCh)
		close(floatCh)
	}()

	for val := range results {
		fmt.Println(val)
	}
}
