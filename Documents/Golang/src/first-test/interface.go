package main

import "fmt"

func greet(name HasName) {
	fmt.Println("Hjai", name.GetName())
}

type HasName interface {
	GetName() string
}

type Orang struct {
	name string
}

func (o Orang) GetName() string {
	return o.name
}

////

func legsCount(legs HasLegs) {
	fmt.Println("legs count :", legs.getLegs())
}

type HasLegs interface {
	getLegs() int
}

type Animal struct {
	legs int
}

func (a Animal) getLegs() int {
	return a.legs
}

func main() {
	var indra Orang
	indra.name = "Indra"
	greet(indra)

	var monki Animal
	monki.legs = 4
	legsCount(monki)

}
