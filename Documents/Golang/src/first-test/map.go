package main

import "fmt"

func main() {
	person := make(map[string]string){
		"name": "Indra",
		"age":  "18",
	}
	fmt.Println(person(len))

	var orang = make(map[string]string)
	orang["nik"] = "21211"
	orang["age"] = "12"

	delete(orang, "age")

	fmt.Println(len(orang))

}
