// 問題
// 文字列 S , T が与えられるので、 T が S の部分文字列かどうかを判定してください。
// なお、 S の部分文字列とは、 S の連続した 1 文字以上を取り出してできる文字列のことです。

// strings.Contains(s, substr string) boolを使うと良さそう

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	const (
		Yes = "YES"
		No  = "NO"
	)

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Fprintf(os.Stdin, "scan err")
		os.Exit(1)
	}
	fullText := scanner.Text()
	if !scanner.Scan() {
		fmt.Fprintf(os.Stdin, "scan err")
		os.Exit(1)
	}
	partText := scanner.Text()

	if strings.Contains(fullText, partText) {
		fmt.Println(Yes)
	} else {
		fmt.Println(No)
	}

}
