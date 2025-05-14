// 問題
// 1 行目に整数 N, M が与えられます。
// 2 行目に M 個の整数 a_1, a_2, ..., a_M が与えられます。
// M 個の整数に N が何個あるか数え、出力してください。
// また、N は M 個の整数の中に 1 個以上含まれるものとします。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// ターゲットを取得
	if !scanner.Scan() {
		fmt.Println("scan err")
		os.Exit(1)
	}
	inputs := strings.Fields(scanner.Text())
	target := inputs[0]

	if !scanner.Scan() {
		fmt.Println("scan err")
		os.Exit(1)
	}
	count := 0

	for _, str := range strings.Fields(scanner.Text()) {
		if str == target {
			count++
		}
	}

	fmt.Println(count)
}
