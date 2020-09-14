package main

func main() {
	// 変数を定義
	var n2 int
	// 値を代入する
	n2 = 200
	println(n2)

	// 変数定義+代入
	var n1 int = 100
	println(n1)

	// 型を省略(int型になる)
	// ここが型推論？
	var n3 = 300
	println(n3)

	// varを省略
	// 関数内のみでしかできない
	n4 := 400
	println(n4)

	// まとめる
	var (
		n = 500
		m = 600
	)
	println(n)
	println(m)
}
