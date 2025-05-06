// 問題
// 1 行で文字列 S が与えられるので、 S が "paiza" である時は "YES", そうでない時は "NO" を出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const (
		Yes = "YES"
		No  = "NO"
	)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	if text == "paiza" {
		fmt.Println(Yes)
	} else {
		fmt.Println(No)
	}
}
