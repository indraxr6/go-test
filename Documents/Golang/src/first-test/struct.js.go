package main

import "fmt"

type person struct {
	name, addr string
	age, nik   int
	married    bool
}

func main() {

	//method 1
	var indra person
	indra.name = "ind wbw"
	indra.addr = "mlg"
	indra.nik = 6969
	indra.age = 21
	indra.married = true

	fmt.Println(indra)

	//method 2 (standard)
	zaim := person{
		name:    "zaim",
		addr:    "mlg",
		age:     18,
		married: true,
	}
	fmt.Println(zaim)

	//method 3
	paiz := person{"paiz", "blitar", 12, 6969, false}
	fmt.Println(paiz)

}
