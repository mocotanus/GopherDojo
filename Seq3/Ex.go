package main

func main() {
	// 【TRY】スライスの利用
	// n1 := 19
	// n2 := 86
	// n3 := 1
	// n4 := 12

	// sum := n1 + n2 + n3 + n4
	// println(sum) // 118

	// ↓3つの変数しか使わないように修正
	ns := []int{19, 86, 1, 12}
	var sum int
	for i := 0; i < len(ns); i++ {
		sum += ns[i]
	}
	println(sum)

}
