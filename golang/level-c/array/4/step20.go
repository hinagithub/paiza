// 問題
// 1 行目に整数 N, M が与えられます。
// 2 行目に M 個の文字列 s_1, s_2, ..., s_M が半角スペース区切りで与えられます。
// N 番目の文字列 s_N を出力してください。

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
	scanner.Scan()
	inputs := strings.Fields(scanner.Text())
	target, err := strconv.Atoi(inputs[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "atoi err \n %d", err)
		os.Exit(1)
	}
	scanner.Scan()

	strs := strings.Fields(scanner.Text())
	fmt.Println(strs[target-1])
}
