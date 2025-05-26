// 問題
// 複数の文字列があります。文字列の数を出力してください。

// Hello
// paiza
// 1234
// pa13

package main

import "fmt"

func main() {
	words := []string{
		"Hello",
		"paiza",
		"1234",
		"pa13",
	}
	fmt.Println(len(words))
}
