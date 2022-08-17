package main

import "fmt"

func main() {

	// Arr := []int{1,2,3,5,6}

	// fmt.Println(Arr)
	
	var month = [14]string {
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
		"Jul",
		"Aug",
		"Sep",
		"Oct",
		"Nov",
		"Des",
	}
	var sliced = append(month[2:])

	// sliced[0] = "Mar Changed"
	// fmt.Println(sliced)

	// var slice2 = month[10:]

	// var slice3 = append(month, "Recently Added")
	sliced[4] = "wkwkwk"


	fmt.Println(sliced)

	




	
}

 