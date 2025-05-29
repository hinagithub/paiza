// 問題
// 文字列Sが与えられます。Sがpaizaと一致する場合はYESを、
// 一致しない場合はNOを出力してください。

package main

import "fmt"

func main() {

	const (
		TargetWord = "paiza"
		Yes        = "YES"
		No         = "NO"
	)

	var text string
	fmt.Scan(&text)

	if text == TargetWord {
		fmt.Println(Yes)
	} else {
		fmt.Println(No)
	}
}
