package main

import "fmt"

// empty interface accepts all data type
func any(x int) interface{} {
	if x == 1 {
		return "ok"
	} else if x == 2 {
		return true
	} else {
		return 69
	}
}

func main() {
	test := any(31831)
	fmt.Println(test)
}
