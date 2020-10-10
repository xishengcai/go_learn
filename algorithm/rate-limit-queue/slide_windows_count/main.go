package main

import (
	"container/list"
	"fmt"
	"time"
)

type part struct {
	Time  int64
	Count int
}

func main() {
	var listPart = list.New()
	nowSecond := time.Now().Unix()
	if listPart.Len() == 0 {
		for i := 0; i < 10; i++ {
			listPart.PushBack(part{
				Time:  nowSecond + int64(i),
				Count: i,
			})
		}
	}

	go func() {
		for {
			nowSecond = time.Now().Unix()
			time.Sleep(time.Second * 1)
			doSomething("remove")
			listPart.Remove(listPart.Front())
			listPart.PushBack(part{
				Time:  nowSecond + 1,
				Count: 0,
			})
		}

	}()

	for {
		nowSecond = time.Now().Unix()

		if listSum(listPart) >= 500 {
			fmt.Println("reach limit")
			time.Sleep(time.Second*1)
			continue
		}

		// 打印方便调试
		for e := listPart.Front(); e != nil; e = e.Next() {
			fmt.Printf("e = %v\n", e.Value)
		}

		time.Sleep(time.Millisecond * 10)
		doSomething("+")
		partElement := listPart.Back().Value.(part)
		listPart.Back().Value = part{
			Time:  partElement.Time,
			Count: partElement.Count + 1,
		}
	}
}

func doSomething(content string) {
	fmt.Println("do something ", content)
}

// 队列求和
func listSum(l *list.List) int {
	var sum int
	for e := l.Front(); e != nil; e = e.Next() {
		sum += e.Value.(part).Count
	}
	return sum
}
