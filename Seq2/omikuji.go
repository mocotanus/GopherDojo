package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// ランダムな値の作り方
	// https://play.golang.org/p/YJbDjyetOxP

	// // 現在時刻を数値で取得する
	// t := time.Now().UnixNano()
	// // 乱数のたねを設定
	// rand.Seed(t)
	// // xは0-5の間の値になる(6個の6)
	// s := rand.Intn(6) + 1
	// println(s)

	// さいころを振る
	rand.Seed(time.Now().UnixNano())
	me := rand.Intn(6) + 1
	switch me {
	case 6:
		fmt.Println("大吉")
	case 5, 4:
		fmt.Println("中吉")
	case 3, 2:
		fmt.Println("吉")
	case 1:
		fmt.Println("凶")
	default:
		fmt.Println("")
	}
}
