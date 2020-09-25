package main

// 偶数かどうか
func isEven(i int) bool {
	return i%2 == 0
}

func main() {
	// 【TRY】奇数偶数判定関数
	for i := 1; i <= 100; i++ {
		print(i)
		// if i%2 == 0 {
		if isEven(i) {
			println("-偶数")
		} else {
			println("-奇数")
		}
	}

}
