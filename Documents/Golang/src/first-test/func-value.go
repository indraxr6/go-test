package main

import "fmt"

func sayBye(name int) int {
	return 12 + name
}

func main() {
	gBye := sayBye
	fmt.Println(gBye(13))

}
