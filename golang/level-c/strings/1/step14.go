// 問題
// 大文字と小文字のアルファベットが混ざった文字列 S が与えられるので、S を全て大文字にした文字列を出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Fprintf(os.Stderr, "scan err")
		os.Exit(1)
	}
	text := scanner.Text()
	fmt.Println(strings.ToUpper(text))
}
