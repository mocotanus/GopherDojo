package main

// 【TRY】複数戻り値の利用
// 値を入れ替えるswap関数を実装してください
func swap(x int, y int) (int, int) {
	return y, x
}

func main() {
	n, m := swap(10, 20)
	println(n, m)
}
