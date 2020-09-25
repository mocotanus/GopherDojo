package main

import "fmt"

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

type T int

func (t *T) f() { println("hi") }

func main() {
	// 100をHex型として代入
	var hex Hex = 100
	// Stringメソッドを呼び出す
	fmt.Println(hex.String())

	var v T
	(&v).f()
	v.f()
}
