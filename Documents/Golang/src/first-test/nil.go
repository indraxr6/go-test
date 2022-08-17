package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var orang map[string]string = newMap("who??")
	if orang == nil {
		fmt.Println("empty field")
	} else {
		fmt.Println(orang)
	}

}
