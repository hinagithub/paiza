// 問題
// 複数の文字列があります。すべての文字列を改行区切りで出力してください。

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

	for _, str := range strs {
		fmt.Println(str)
	}
}
