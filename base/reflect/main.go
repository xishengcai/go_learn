package main

import (
	"fmt"
	"github.com/codegangsta/inject"
)

type SpecialString interface{}

func main() {
	fmt.Println(inject.InterfaceOf((*interface{})(nil)))
	fmt.Println(inject.InterfaceOf((*SpecialString)(nil)))
}
