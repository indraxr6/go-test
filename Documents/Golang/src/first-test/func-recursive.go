package main

import "fmt"

func factorialLoop(val int) int {
	res := 1
	for i := val; i > 0; i-- {
		res *= i
	}
	return res
}

func factorialRecursive(val int) int {
	if val == 1 {
		return 1
	} else {
		return val * factorialRecursive(val-1)
	}
}

func main() {
	//result := factorialLoop(5)
	//fmt.Println(result)
	res := factorialRecursive(4)
	fmt.Println(res)

}
