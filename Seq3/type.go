package main

import "fmt"

func main() {
	// 【TRY】組み込み型（数値）
	var sum int
	sum = 5 + 6 + 3
	avg := sum / 3 // avgには4が入っている
	// avg := float64(sum) / 3.0
	println(avg)

	// コンパイルエラー：constant 4.5 truncated to integer
	//4.5が整数じゃないため

	// 4.5の型を調べる
	fmt.Printf("%T\n", 4.5) // 4.5の型はfloat64

	// if avg > 4.5 {
	if float64(avg) > 4.5 {
		println("good")
	}

	// 型の違う値は計算できないらしい

}
