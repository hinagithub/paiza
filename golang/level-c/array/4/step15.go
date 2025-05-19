// 問題
// 複数の文字列があります。すべての文字列を改行区切りで出力してください。

// eight
// one
// three
// paiza
// pa13
// 813

package main

import "fmt"

func main() {
	strs := []string{
		"eight",
		"one",
		"three",
		"paiza",
		"pa13",
		"813",
	}
	for _, str := range strs {
		fmt.Println(str)
	}
}
