package main

import "fmt"

// varで宣言する場合は関数の外から呼び出せる
var from_outside string = "test"

func main() {
	// varを使って変数を宣言(型も指定しないといけない)
	var i int = 1
	var f64 float64 = 1.4
	var s string = "variable"
	var t, f bool = true, false
	fmt.Println(i, f64, s, t, f)

	// :=で宣言する場合は型指定はいらない
	vi := 1
	vf64 := 2.3

	fmt.Println(vi, vf64)

	// 関数の外にあるvarで宣言された変数
	fmt.Println(from_outside)
}
