package main

import (
	"errors"
	"fmt"
)

func divide(val int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("dividing with 0")
	} else {
		result := val / divider
		return result, nil
	}
}

func main() {
	res, err := divide(10, 0)
	if err == nil {
		fmt.Println("Result ", res)
	} else {
		fmt.Println("Error", err.Error())
	}
}
