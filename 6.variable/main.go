package main

import "fmt"

func main() {
	// varを使って変数を宣言(型も指定しないといけない)
	var i int = 1
	var f64 float64 = 1.4
	var s string = "variable"
	var t, f bool = true, false
	fmt.Println(i, f64, s, t, f)
}
