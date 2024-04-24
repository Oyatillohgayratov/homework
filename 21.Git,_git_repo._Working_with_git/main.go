package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu   sync.Mutex
	data map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (sm *SafeMap) Get(key string) int {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return sm.data[key]
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()	
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

func main() {
	safeMap := NewSafeMap()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		safeMap.Set("key1", 1)
		safeMap.Set("key2", 2)
		safeMap.Set("key3", 3)
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Value of key1:", safeMap.Get("key1"))
		fmt.Println("Value of key2:", safeMap.Get("key2"))
		fmt.Println("Value of key3:", safeMap.Get("key3"))
	}()

	wg.Wait()

	safeMap.Delete("key2")
	fmt.Println("Value of key2 after deletion:", safeMap.Get("key2"))
}
