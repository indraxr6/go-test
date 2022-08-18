package main

import "fmt"

type People struct {
	Name string
	Age  int
}

func (name *People) Chad() {
	name.Name = "Bro " + name.Name
}

func (age *People) ChadAge() {
	age.Age = 18
}

func main() {
	bin := People{"Inibin", 12}
	bin.Chad()
	bin.ChadAge()
	fmt.Println(bin.Name)
	fmt.Println(bin.Age)

}
