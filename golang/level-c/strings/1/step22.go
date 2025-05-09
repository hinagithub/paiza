// 文字列 S が与えられるので、 S を整数に変換できる場合には "YES" , そうでない場合は "NO" を出力してください。
// なお、文字列 S を数値に変換できるとは、S の全ての文字が
// { 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 } のいずれかであることをいいます。

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	const (
		Yes = "YES"
		No  = "NO"
	)

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "scan err")
		os.Exit(1)
	}
	text := scanner.Text()

	for _, r := range text {
		if !unicode.IsDigit(r) {
			fmt.Println(No)
			return
		}
	}
	fmt.Println(Yes)
}
