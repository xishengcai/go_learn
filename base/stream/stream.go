package main

import (
	"fmt"
	"xishengcai/stream"
)

func main() {
	type student struct {
		Name string
		Age  int
	}
	stream.New([]student{
		{"a1", 1},
		{"a2", 7},
		{"a3", 0},
		{"a4", 10},
		{"a5", 5},
	}, false).Map(func(v interface{}) interface{} {
		s := v.(student)
		s.Age += 10
		return s
	}).Sorted(func(i, j interface{}) bool {
		s1 := i.(student)
		s2 := j.(student)
		return s1.Age < s2.Age
	}).ForEach(func(v interface{}) {
		fmt.Println(v)
	})

}
