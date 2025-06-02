// 問題
// 長さ N の数列 a (a_1, a_2, ..., a_N) が与えられます。
// この数列の中に 1 が何個あるか出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	readline(scanner) // 1行目の要素数nは使用しないのでスキップ
	strNumbers := strings.Fields(readline(scanner))
	count := 0
	for _, strNum := range strNumbers {
		if strNum == "1" {
			count++
		}
	}
	fmt.Println(count)
}

func readline(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "failed to scan input")
		os.Exit(1)
	}
	return scanner.Text()
}
