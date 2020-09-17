package main

import "strconv"

func main() {

	println("(･ω･)if文を使った場合~~~~~~~~~~")
	for i := 1; i <= 100; i = i + 1 {
		str := ""
		if i%2 == 0 {
			str = "偶数"
		} else {
			str = "奇数"
		}

		// println(string(i) + "-")// A,Bになってしまった
		// 整数値は Unicode コードポイントとして解釈されるとのこと
		//https://www.delftstack.com/ja/howto/go/how-to-convert-an-int-value-to-string-in-go/

		// 文字列変換関係は「strconv」パッケージのItoa 関数を使えばいい
		println(strconv.Itoa(i) + "-" + str)
	}

	println("")
	println("(･ω･)switchを使った場合~~~~~~~~~~")
	for i := 1; i <= 100; i = i + 1 {
		str := ""
		switch {
		case i%2 == 0:
			str = "偶数"
		default:
			str = "奇数"
		}

		println(strconv.Itoa(i) + "-" + str)
	}
}
