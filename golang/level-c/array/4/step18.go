// 問題
// 整数 N が与えられます。
// 以下の文字列のうち、N 番目の文字列を出力してください。

// good
// morning
// paiza
// 813
// pa13

package main

import "fmt"

func main() {

	strs := []string{
		"good",
		"morning",
		"paiza",
		"813",
		"pa13",
	}

	var n int
	fmt.Scan(&n)

	fmt.Println(strs[n-1])

}
