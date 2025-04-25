// 問題
// 文字列one two three four fiveを、半角スペースで分割して出力してください。

package main

import (
	"fmt"
	"strings"
)

func main() {
	const sentence string = "one two three four five"
	texts := strings.Split(sentence, " ")
	for _, text := range texts {
		fmt.Println(text)
	}
}
