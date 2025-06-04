// 問題
// N 個の整数 M_1, M_2, ..., M_N があります。
// i 番目の M を M_i とするとき、M_i * i を改行区切りで出力してください。
// 例えば、M_5 が 3 の場合、3 * 5 = 15 となります。

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
	readLine(scanner) // 1行目の要素数は使用しないのでスキップ
	inputs := strings.Fields(readLine(scanner))

	for i, input := range inputs {
		num := atoi(input)
		fmt.Println(num * (i + 1))
	}

}

func readLine(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "failed to scan input")
		os.Exit(1)
	}
	return scanner.Text()
}

func atoi(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to atoi: %v \n", err)
		os.Exit(1)
	}
	return num
}
