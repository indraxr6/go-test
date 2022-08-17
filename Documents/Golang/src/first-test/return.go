package main

import "fmt"

//func chain(name string) string {
//	if name == "" {
//		return "iyh kosong"
//	} else {
//		return "ada"
//	}
//}

//func main() {
//	//	//name := chain("bin")
//	//	name := chain("dsadsa")
//	//	fmt.Println(name)
//	//}

func chain() (string, int) {
	return "bin", 18

}

func main() {
	first, _ := chain()
	fmt.Println(first)
}
