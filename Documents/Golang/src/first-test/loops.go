package main

import "fmt"

func main() {

	person := make(map[string]string)
	person["name"] = "1221"
	person["address"] = "83883"

	for key, value := range person {
		fmt.Println(key, "= ", value)
	}

	//1 loops

	// val := 0

	// for val := 0; val <= 15; val++ {
	// 	fmt.Println(val)

	// }

	//2 array loops

	// slice := []string{"Indra", "Faiz", "Zaim"}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }

	//3 array loops

	// slice := []string{"Indra", "Faiz", "Zaim"}

	// for _, value := range slice {
	// 	fmt.Println(value)
	// }
}
