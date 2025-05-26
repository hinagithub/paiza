// 問題
// 1 行目に整数 N, M, L が与えられます。
// 2 行目に M 個の文字列 s_1, s_2, ..., s_M が半角スペース区切りで与えられます。
// N 番目の文字列 s_N の L 番目の文字を出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// ターゲット要素番号を取得
	scanner.Scan()
	inputs := strings.Fields(scanner.Text())
	targetWordIndex, _ := strconv.Atoi(inputs[0])
	targetCharIndex, _ := strconv.Atoi(inputs[2])

	// 文字の配列を取得
	scanner.Scan()
	words := strings.Fields(scanner.Text())
	fmt.Println(string(words[targetWordIndex-1][targetCharIndex-1]))
}
