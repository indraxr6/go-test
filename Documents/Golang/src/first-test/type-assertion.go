package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	var res = random()
	switch value := res.(type) {
	case int:
		fmt.Println("int", value)
	case string:
		fmt.Println("string", value)
	case interface{}:
		fmt.Println("whatever / any", value)
	default:
		fmt.Println("???")
	}
}
