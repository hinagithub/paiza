// 問題
// 整数 A, B, C が与えられます。以下のアルゴリズムを実行してください。
// 変数 N を 0 で初期化する
// N に A を足した値を N へ代入する
// N に B をかけた値を N へ代入する
// N を C で割ったあまりを N へ代入する
// N を出力する

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
	if !scanner.Scan() {
		fmt.Println(os.Stderr, "scan err")
		os.Exit(1)
	}
	inputs := strings.Fields(scanner.Text())

	a, _ := strconv.Atoi(inputs[0])
	b, _ := strconv.Atoi(inputs[1])
	c, _ := strconv.Atoi(inputs[2])

	n := 0
	n = n + a
	n = n * b
	n = n % c
	fmt.Println(n)
}
