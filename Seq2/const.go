package main

import "fmt"

func main() {
	// 定数

	// ◼️名前付き定数
	// 型のある定数
	const c1 int = 100
	println(c1)

	// 型のない定数
	const c2 = 200
	println(c2)

	// 定数式
	const c3 = "Hello, " + "世界"
	println(c3)

	// まとめる書き方
	const (
		x = 1000
		y = 2000
	)
	println(x)
	println(y)

	// 右辺の省略
	// グループ化した名前付き定数で使用できる
	// 省略した場合、１つめの右辺と同じ値が入る
	const (
		a = 1 + 2
		b
		c
	)
	println(a)
	println(b)
	println(c)

	println("iota")
	// iota
	// グプープかされた名前付き定数の定義で使用
	// 0から始まり、1ずつ加算される値
	const (
		i1 = iota
		// 省略できる
		i2
	)
	const (
		// 式の中でも使える
		i3 = 1 << iota
		i4
		i5
	)
	// println(i1)
	// println(i2)
	// println(i3)
	// println(i4)
	// println(i5)
	fmt.Println(i1, i2, i3, i4, i5)

}
