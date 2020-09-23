package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// ●コンポジット型
	// コンポジット型を要素として持つコンポジット型
	// 要素も型リテラルで書かれているだけ
	// スライスの要素がスライスの場合（2次元スライス）
	var t1 [][]int
	fmt.Println(t1)
	// マップの値がスライスの場合
	var t2 map[string][]int
	fmt.Println(t2)
	// 構造体のフィールドの型が構造体
	var t3 struct {
		A struct {
			N int
		}
	}
	fmt.Println(t3)

	// -------------------------------------
	// ●ユーザ定義型
	// typeで名前を付けて新しい型を定義する
	// type 型名 基底型

	// 組み込み型を基にする
	type MyInt int
	// 他のパッケージの型を基にする
	type MyWriter io.Writer
	// 型リテラルを基にする
	type Person struct {
		Name string
	}

	// キャスト
	type MyInt2 int
	var n int = 100
	m := MyInt2(n)
	n = int(m)
	fmt.Println(n, m)

	// デフォルトの型からユーザ定義型へキャストできる場合
	// 10秒を表す（time.Duration型、type Duration int64）
	d := 10 * time.Second // time.Secondはtime.Duration型(type Duration int64)
	fmt.Println(d)

	// 型エイリアス（Go 1.9以上）を定義する
	// 	完全に同じ型
	// キャスト不要
	// エイリアスの方ではメソッド定義はできない
	type Applicant = http.Client
	// 型名を出力する%Tが同じ元の型名を出す
	fmt.Printf("%T\n", Applicant{}) // http.Client

}
