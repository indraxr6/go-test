package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("i([a-n])dra")
	fmt.Println(regex.MatchString("izdra"))

	src := regex.FindAllString("indra izra itdra", -1)
	fmt.Println(src)
}
