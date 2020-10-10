package main

import (
	"fmt"
	"time"
)

func main() {
	goroutine()
}


func normal(){
	var rate = 10        // 令牌生成速度，每秒10个
	var tokens int64 = 0 // 当前令牌数量
	var timeStamp = time.Now().Unix()

	for {
		now := time.Now().Unix()
		di := tokens + (now-timeStamp)*int64(rate) //计算过去的时间段产生的令牌数
		if di > 0 {
			tokens = di
		} else {
			tokens = 0
		}
		timeStamp = now
		if tokens < 1 {
			fmt.Println("reach limit")
			time.Sleep(500 * time.Millisecond)
		} else {
			tokens--
			doSomething()
		}
	}
}

func goroutine() {
	var ch = make(chan int64, 100)
	go func() {
		// 开启协程，每隔100ms往管道里面放一个数，也就是实现每秒10s的限速效果
		var ticker = time.NewTicker(1 * time.Millisecond)
		for {
			select {
			case <-ticker.C:
				ch <- 1
			}
		}
	}()
	for {
		<-ch //利用管道的阻塞特性
		doSomething()
	}
}

func doSomething() {
	fmt.Printf(" %s :do something \n", time.Now().Format("2006-01-02 15:04:05"))
}
