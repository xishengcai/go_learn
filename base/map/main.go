package main

import (
	"fmt"
	"strings"
)

type aaa struct {
	A string
	*BB
}

type BB struct {
	Name string
}

func main() {

	m := map[int]string{}

	m[1] = "1"

	fmt.Print("2  ", m[2])
	x, ok := m[2]
	if ok {
		fmt.Print(x)
	}

	o := "`"
	y := fmt.Sprintf(`%s hello %s`, o, o)
	fmt.Println(y)

	repo := "https://xishengcai/cloud.git"
	repoNameList := strings.Split(repo, "/")
	length := len(repoNameList)
	org := strings.Join(repoNameList[3:length-1], "/")
	repoName := strings.Split(repoNameList[length-1], ".git")[0]
	fmt.Println(org)
	fmt.Println(repoName)

}
