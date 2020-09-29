package main

import "fmt"

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

func main() {
	// 明示的には呼び出せない
	// init()

	fmt.Println("main")
}
