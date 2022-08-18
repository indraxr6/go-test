package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("indra")
	data.PushBack("wibowo")
	data.PushBack("anjir")
	data.PushFront("first")

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
	//fmt.Println(data.Front().Value)
	//fmt.Println(data.Back().Value)

}
