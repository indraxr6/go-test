package main

import "fmt"

func sumAll(num ...int) int {
	total := 0
	for _, number := range num {
		total += number
	}
	return total
}

func main() {
	//slice := []int{10, 20, 30, 40, 50}
	//total := sumAll(slice...)
	//
	//fmt.Println(total)

	summing := sumAll()
	fmt.Println(summing)

}
