package main

import "fmt"

func main() {
	// マップはキーとバリュー(値)
	// キーには「==」で比較できる型しかNG

	// キーと値を指定する
	// var m map[string]int

	// マップの初期化
	// ゼロ値はnil
	var m map[string]int
	fmt.Println(m)
	// makeで初期化
	m = make(map[string]int)
	fmt.Println(m)
	// 容量を指定できる
	m = make(map[string]int, 10)
	fmt.Println(m)
	// リテラルで初期化
	m2 := map[string]int{"x": 10, "y": 20}
	fmt.Println(m2)
	// 空の場合
	m3 := map[string]int{}
	fmt.Println(m3)
	fmt.Println("----------")

	// マップの操作
	m4 := map[string]int{"x": 10, "y": 20}
	// キーを指定してアクセス
	println(m4["x"])
	// キーを指定して入力
	m4["z"] = 30
	// 存在を確認する
	n, ok := m4["z"]
	println(n, ok)
	// キーを指定して削除する
	delete(m4, "z")
	// 削除されていることを確認
	n, ok = m4["z"] // ゼロ値とfalseを返す
	println(n, ok)

}
