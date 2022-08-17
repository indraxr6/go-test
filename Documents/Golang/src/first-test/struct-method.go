package main

import "fmt"

type Person struct {
	name, addr string
	//age, nik   int
	//married    bool
}

func (data Person) greet() {
	fmt.Println("Henlo", data.name)
	fmt.Println("you're livin in", data.addr)
}

func main() {
	indra := Person{name: "indra", addr: "malang"}
	indra.greet()

}
