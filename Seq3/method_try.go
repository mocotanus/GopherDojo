package main

type MyInt int

// func (n MyInt) Inc() { n++ }
// ↑だとコピーになるので、レシーバ(呼び出し元)の値は変更されない
func (n *MyInt) Inc() { *n++ }

// ↑ポインタ型にすると、レシーバへの変更を呼び出し元に伝えることができる

// 【TRY】レシーバに変更を与える
// ポインタ型にする必要がある
func main() {
	var n MyInt
	println(n)

	// レシーバがポインタの場合もドットでアクセスする
	n.Inc()
	println(n)

	// &をつけてもつけなくても同じ意味
	(&n).Inc()
	println(n)

}
