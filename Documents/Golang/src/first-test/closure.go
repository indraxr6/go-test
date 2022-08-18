package main

import (
	"fmt"
)

func main() {
	name := "indra"
	counter := 0
	// scope function bisa mengakses data atasnya, namun tidak bisa mengakses data dibawahnya
	increment := func() {
		name = "indrrr"
		fmt.Println("add")
		counter++
	}
	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
