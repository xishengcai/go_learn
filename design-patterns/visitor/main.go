package main

import (
	"fmt"
)

type Room interface {
	Accept(Visitor)
}

type Visitor interface {
	VisitBuilding(building Building)
	VisitApartment(apartment Apartment)
}

type Apartment struct {
	Name string
}

func (a Apartment) Accept(v Visitor) {
	v.VisitApartment(a)
}

type Building struct {
	Name string
}

func (b Building) Accept(v Visitor) {
	v.VisitBuilding(b)
}

type Cleaner struct{}

func (c Cleaner) VisitApartment(a Apartment) {
	fmt.Printf("cleanup Apartment %s\n", a.Name)
}

func (c Cleaner) VisitBuilding(b Building) {
	fmt.Printf("cleanup Building %s\n", b.Name)
}

func main() {

	targets := []Room{
		Apartment{Name: "101"},
		Apartment{Name: "103"},
		Apartment{Name: "201"},
		Building{Name: "B201"},
		Building{Name: "B201"},
		Building{Name: "B201"},
	}

	c := Cleaner{}
	for _, item := range targets {
		item.Accept(c)
	}
}
