// 文字列He likes paizaを、半角スペースで分割して出力してください。

package main

import (
	"fmt"
	"strings"
)

func main() {
	const sentence string = "He likes paiza"
	texts := strings.Split(sentence, " ")
	for _, text := range texts {
		fmt.Println(text)
	}
}
