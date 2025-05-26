// 問題
// 1 行目に整数 A, B, N が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// N 個の整数のうち、その数が A だった場合、B に書き換えてください。
// 書き換えた N 個の整数を改行区切りで出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// ターゲットと置換文字を取得
	inputs := strings.Fields(readline(scanner))
	target := inputs[0]
	replacement := inputs[1]

	// 出力
	for _, str := range strings.Fields(readline(scanner)) {
		if str == target {
			fmt.Println(replacement)
		} else {
			fmt.Println(str)
		}
	}

}

func readline(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Println("scan err")
		os.Exit(1)
	}
	return scanner.Text()
}
