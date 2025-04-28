// 問題
// N 個の文字列 S_1, ... , S_N と、Q 個の文字列 T_1, ... , T_Q が与えられます。
// 各 T_i について、以下の処理を行ってください。
// ・ S_j == T_i を満たす最小の j を出力する。
// ただし、そのような j が存在しない場合は -1 を出力する。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// 最初の行で N と Q を読み取る
	sc.Scan()
	inputs := strings.Fields(sc.Text())
	n, _ := strconv.Atoi(inputs[0])

	// S の文字列を読み込む
	texts := make([]string, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		texts[i] = sc.Text()
	}

	// Q 回のクエリを処理
	for sc.Scan() {
		target := sc.Text()
		result := findIndex(texts, target)
		fmt.Println(result)
	}
}

func findIndex(texts []string, target string) int {
	result := -1
	for i, text := range texts {
		if text == target {
			result = i + 1
			return result
		}
	}
	return result
}
