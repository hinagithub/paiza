// 問題
// 1 行目に整数 N, M が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// N 個の整数の中に、整数 M が含まれているなら Yes、含まれていないなら No を出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const (
		Yes = "Yes"
		No  = "No"
	)
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("scan err")
		os.Exit(1)
	}

	// 探索する文字列を取得
	inputs := strings.Fields(scanner.Text())
	target := inputs[1]

	// 数列を取得
	if !scanner.Scan() {
		fmt.Println("scan err")
		os.Exit(1)
	}
	strings := strings.Fields(scanner.Text())

	// 探索
	found := false
	for _, str := range strings {
		if str == target {
			found = true
			break
		}
	}

	if found {
		fmt.Println(Yes)
	} else {
		fmt.Println(No)
	}

}
