package main

import (
	"fmt"
	"time"
)

func main() {
	var capacity int64 = 100 // 桶容量
	var rate = 10            // 出水速度，每秒10个
	var water int64 = 0      // 当前水量
	var lastTime = time.Now().Unix()

	for {
		now := time.Now().Unix()
		di := water - (now-lastTime)*int64(rate) //计算过去的时间内产生的水量
		if di > 0 {
			water = di
		} else {
			water = 0
		}
		lastTime = now
		if water < capacity {
			water++
			doSomething(water)
		} else {
			fmt.Println("reach limit")
			time.Sleep(3 * time.Second)
		}
	}
}

func doSomething(water int64) {
	fmt.Println("do something ", water)
}
