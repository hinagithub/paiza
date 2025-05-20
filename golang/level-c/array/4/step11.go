// 問題
// 1 行目に整数 N, M, K が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// 左から M 番目に K を挿入し、挿入した後の N + 1 個の整数を改行区切りで出力してください。
// なお、左端を 1 番目とします。

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
		fmt.Println("scan error")
		os.Exit(1)
	}
	inputs := strings.Fields(scanner.Text())
	insertIndex, err := strconv.Atoi(inputs[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "atoi err: %d", err)
		os.Exit(1)
	}
	insertValue := inputs[2]

	if !scanner.Scan() {
		fmt.Println("scan error")
		os.Exit(1)
	}
	for i, strNum := range strings.Fields(scanner.Text()) {
		if i == insertIndex-1 {
			fmt.Println(insertValue)
		}
		fmt.Println(strNum)

	}
}
