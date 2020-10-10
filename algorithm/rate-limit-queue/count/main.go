package main

import (
	"fmt"
	"time"
)

func main() {
	var count = 1
	var sTime = time.Now()

	for {
		if time.Now().Sub(sTime).Seconds() > 1 {
			count = 1
			sTime = time.Now()
		}
		if count > 100  {
			time.Sleep(time.Millisecond * 100)
			fmt.Println("reach limit...")
			continue
		}
		doSomething(count)
		count++
	}
}

func doSomething(n int) {
	fmt.Printf("do something #%d\n", n)
}
