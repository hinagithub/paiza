// 問題
// 文字列 S が与えられるので、 S の前後を反転させた文字列を出力してください。
// 例えば、 S = "abcde" のとき、前後を反転させた文字列は "edcba" となります。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "scan err")
		os.Exit(1)
	}
	text := []rune(scanner.Text())
	results := make([]rune, len(text))

	for i := len(text) - 1; i >= 0; i-- {
		results[len(text)-1-i] = text[i]
	}
	fmt.Println(string(results))
}
