package main

import "fmt"

// 【TRY】ポインタ
// 値を入れ替えるswap2関数を実装してください

func swap2(x *int, y *int) {
	// * はポインタ型変数から値型変数を取得する演算子
	// *でポインタの指す先に値を入れる
	// 例)*x = 100

	// tmp := *y
	// *x = tmp

	// tmp = *x// 置き換わった後の値が入ってしまうのでNG(´・ω・`)
	// *y = tmp

	tmp := *x

	*x = *y
	*y = tmp
}

// ポインタを使う練習
func pointerSample() {
	var tmp string = "tmp"

	// ポインタを取得
	// & は値型変数からポインタ(メモリのアドレス)型変数を取得する演算子
	tmpPointer := &tmp
	// 値を取得する
	// * はポインタ型変数から値型変数を取得する演算子
	tmpValue := *tmpPointer

	fmt.Println(tmpPointer) // 0xc000010200
	fmt.Println(tmpValue)   // tmp
}

func main() {
	// pointerSample()

	n, m := 10, 20
	// &はポインタ(メモリのアドレス)型変数を取得する演算子
	swap2(&n, &m) // &でポインタを取得する
	println(n, m)

}
