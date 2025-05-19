// 問題
// 1 行目に整数 N が与えられます。
// 2 行目に N 個の文字列 s_1, s_2, ..., s_Nが半角スペース区切りで与えられます。
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
	scanner.Scan() // 1行目は使わないのでスキップ
	if scanner.Scan() {
		strs := strings.Fields(scanner.Text())
		for _, str := range strs {
			fmt.Println(str)
		}
	}
}
