// 問題
// 大文字または小文字のアルファベットCが与えられます。
// Cが大文字の場合はYESを、そうではない場合はNOを出力してください。

package main

import (
	"fmt"
	"unicode"
)

func main() {
	var char string
	fmt.Scan(&char)
	if unicode.IsUpper([]rune(char)[0]) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
