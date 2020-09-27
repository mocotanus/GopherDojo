package main

import (
	"fmt"

	"./greeting"
)

func main() {
	// 【TRY】パッケージを分けてみよう
	greeting.Do()

	// fmt.Println(greeting.Do())
	fmt.Println(greeting.Do2())
}
