package main

import "fmt"

type Filter func(string) string

func sayHello(name string, filter Filter) {
	fmt.Println("Hai", filter(name))
}
func sayFilter(name string) string {
	if name == "asu" {
		return "hah??"
	} else {
		return name
	}
}
func main() {
	sayHello("asu", sayFilter)
}
