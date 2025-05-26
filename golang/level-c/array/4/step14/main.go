// 問題
// 複数の文字列があります。上から 3 番目の文字列を出力してください。

// good
// morning
// paiza
// 813
// pa13

package main

import "fmt"

func main() {
	const targetIndex = 3 - 1
	strs := []string{
		"good",
		"morning",
		"paiza",
		"813",
		"pa13",
	}

	fmt.Println(strs[targetIndex])

}
