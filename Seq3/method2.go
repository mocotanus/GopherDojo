package main

import "fmt"

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func main() {
	// 100をHex型として代入
	var hex Hex = 100

	// メソッド値として代入
	f := hex.String
	fmt.Println(f())

	// メソッド式(メソッドを表す式)
	// レシーバを第1引数とした関数になる
	f2 := Hex.String
	fmt.Printf("%T\n%s\n", f2, f2(hex))
}
