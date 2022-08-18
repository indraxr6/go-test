package main

import (
	"fmt"
	"os"
)

func main() {
	hostname, err := os.Hostname()

	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println(err.Error())
	}

	username := os.Getenv("app-usr")
	pw := os.Getenv("app-pw")

	fmt.Println(username)
	fmt.Println(pw)
}
