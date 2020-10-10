package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	limiter := rate.NewLimiter(5, 5) // 速率，容量
	for {
		if limiter.Allow() {
			fmt.Printf("%s: allow\n", time.Now().Format("2006-01-02 15:04:05"))
		}else{
			fmt.Println("reach limit, wait...")
			time.Sleep(time.Millisecond*1000)
		}
	}
}
