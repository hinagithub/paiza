// 問題
// ',' が含まれている文字列 S が与えられるので、S を ',' で分割したときの各要素を出力してください。
// 最後の要素を出力するのを忘れないように注意してください。

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
		fmt.Fprintf(os.Stderr, "sacn err")
		os.Exit(1)
	}
	inputs := strings.Split(scanner.Text(), ",")
	for _, input := range inputs {
		fmt.Println(input)
	}
}
