package main

import "fmt"

func main() {
	// 値を二つ入れれるarray
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	// 宣言した以上の数の値はエラーになる
	//	var b [2]int = [2]int{100, 200}
	//	b = append(b, 300)
	//	fmt.Println(b)

	// sliceは宣言した後に値を追加できる
	var b []int = []int{100, 200}
	b = append(b, 300)
	fmt.Println(b)
}
