package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Database host")
	var user *string = flag.String("user", "root", "Database host")
	var password *string = flag.String("password", "root", "Database host")

	flag.Parse()
	fmt.Println(*host)
	fmt.Println(*user)
	fmt.Println(*password)
}
