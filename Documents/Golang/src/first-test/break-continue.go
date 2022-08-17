package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 4 {
			continue
		}
		fmt.Println(i)

	}

	for i := 0; i < 10; i++ {
		if i == 4 {
			break
		}
		fmt.Println(i)

	}

	//for i := 0; i > 10; i++ {
	//	if i%2 == 0 {
	//		continue
	//	}
	//	fmt.Println(i)
	//
	//}
}
