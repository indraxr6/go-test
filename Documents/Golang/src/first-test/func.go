package main

import "fmt"

func print() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func say(first string, last int) {
	fmt.Println("Woy", first, last)
}

func main() {
	say("indra", 2121)

}
