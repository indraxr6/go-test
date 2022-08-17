package main

import "fmt"

func logging() {
	fmt.Println("log called")
}

func runApp(value int) {
	defer logging()
	fmt.Println("running...")
	result := 10 / value
	fmt.Println("res :", result)
}

func main() {
	runApp(10)
}
