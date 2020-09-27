package greeting

//大文字から始まる関数は外部に公開される。
func Do() {
	println("こんにちは")
}

func Do2() string {
	return "こんにちは"
}
