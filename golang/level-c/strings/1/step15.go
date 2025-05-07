// 問題
// 大文字と小文字のアルファベットが混ざった文字列 S が与えられます。
// S の小文字を全て大文字に、大文字を全て小文字にした文字列を出力してください。

// func strings.Map(mapping func(rune) rune, s string) stringを使うと良さそう

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Fprintf(os.Stderr, "scan err")
		os.Exit(1)
	}
	text := scanner.Text()
	result := strings.Map(swap, text)
	fmt.Println(string(result))
}

func swap(r rune) rune {
	if unicode.IsUpper(r) {
		return unicode.ToLower(r)
	} else {
		return unicode.ToUpper(r)
	}
}
