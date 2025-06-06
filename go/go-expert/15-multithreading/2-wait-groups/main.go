package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(time.Second * 1)
		wg.Done()
	}
}

// Thread 1
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)
	// Thread 2
	go task("A", &waitGroup)
	// Thread 3
	go task("B", &waitGroup)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task anonymous is running\n", i)
			time.Sleep(time.Second * 1)
			waitGroup.Done()
		}
	}()
	// Nothing
	// Exit
	waitGroup.Wait()
}
