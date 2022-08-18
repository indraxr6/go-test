package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("indra", "dr")) // true
	fmt.Println(strings.Split("indra wbw", " ")) // [indra, wbw]
	fmt.Println(strings.ToLower("INDRA"))
	fmt.Println(strings.ToUpper("indra"))
	fmt.Println(strings.ToTitle("indra wbw wkwk"))
	fmt.Println(strings.Trim("indra wbw wkwk", "indra"))
	fmt.Println(strings.ReplaceAll("bin bin oke bin", "bin", "idr"))
	fmt.Println()

}
