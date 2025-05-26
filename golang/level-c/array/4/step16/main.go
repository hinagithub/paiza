// 問題
// 10 個の文字列 s_1, s_2, ..., s_10 が半角スペース区切りで与えられます。
// すべての文字列を改行区切りで出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		for _, str := range strings.Fields(scanner.Text()) {
			fmt.Println(str)
		}
	}
}
