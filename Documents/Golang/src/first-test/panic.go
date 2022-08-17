package main

import "fmt"

func endApp() {
	fmt.Println("app stopped")
	message := recover()
	if message != nil {
		fmt.Println("err message : ", message)
	}
}

func startApp(error bool) {
	defer endApp()
	if error {
		panic("crashed")
	}
	fmt.Println("app running...")
}

func main() {
	startApp(true)

}
