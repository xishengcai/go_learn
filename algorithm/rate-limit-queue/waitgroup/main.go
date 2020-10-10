package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg = sync.WaitGroup{}
	var count = 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go doSomething(&wg)
		count++
		if count > 10 {
			wg.Wait()
			count = 0
		}
	}
}

func doSomething(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 3)
	fmt.Printf(" %s :do something \n", time.Now().Format("2006-01-02 15:04:05"))
}
