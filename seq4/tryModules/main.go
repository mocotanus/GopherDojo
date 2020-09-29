package main

import (
	"fmt"
	"time"

	// "github.com/tenntenn/greeting"
	"github.com/tenntenn/greeting/v2"
)

func main() {
	// ver1
	// fmt.Println(greeting.Do())

	// ver2
	fmt.Println(greeting.Do(time.Now()))


	// ●go.modとは
	// ライブラリのバージョン管理に使うファイル
	// 置き場所は２パターン
	// ・GOPATH配下に作る場合
	// ・GOPATHの外で作る場合

	// 「GOPATHの外で作る場合」についてメモ↓
	// go.modを置きたいディレクトリに移動して、
	// go mod init モジュール名?
	// 例）go mod init seq4

	// go getでライブラリをDLする
	// # 特定のバージョンの2.0.0が欲しい場合
	// go get github.com/tenntenn/greeting/v2@v2.0.0

	// バージョンアップしたい場合
	// go get github.com/tenntenn/greeting/v2@2.1.0

	// @で最新版が取れるらしい
	// go get github.com/tenntenn/greeting/v2@latest

	// 使ってないのを消す
	// go mod tidy


}
