package main

import (
	"fmt"
	"time"
)

func main() {
	clock := time.Now()
	fmt.Println(clock.Day())
	fmt.Println(clock.Year())
	fmt.Println(clock.Month())
	fmt.Println(clock.Second())
	fmt.Println(clock.Nanosecond())

	utc := time.Date(2020, 10, 4, 5, 7, 10, 13, time.UTC)
	fmt.Println(utc)

}
