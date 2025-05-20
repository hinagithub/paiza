// 問題
// 1 行目に整数 N, K が与えられます。
// 2 行目に N 個の文字列 s_1, s_2, ..., s_N が半角スペース区切りで与えられます。
// N 個の文字列を辞書順に並べ替え、K 番目の文字列を出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// ターゲットインデックスを取得
	inputs := strings.Fields(scanner.Text())
	target, err := strconv.Atoi(inputs[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "atoi err %d", err)
		os.Exit(1)
	}

	// 文字列の配列を取得
	scanner.Scan()
	strs := strings.Fields(scanner.Text())

	// 長さチェック
	if target < 0 || target > len(strs)-1 {
		fmt.Printf("target index 「 %d」 out of range.", target)
		os.Exit(1)
	}
	sort.Strings(strs)
	fmt.Println(strs[target-1])

}
