package main

import "fmt"

type Blacklist func(string) bool

func register(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("blocked", name)
	} else {
		fmt.Println("welcum ", name)
	}
}
func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}
	register("indra", blacklist)

	register("admin", blacklist)

}
