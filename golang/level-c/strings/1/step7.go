// 問題
// 与えられる文字列の数 N と、 N 個の文字列が与えられるので、それらを順に末尾に連結した文字列を出力してください。
// 例として、"abc", "def", "ghi" という順で文字列が与えられたとき、連結後の文字列は"abcdefghi" となります。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n, err := strconv.Atoi(readline(scanner))
	if err != nil {
		fmt.Fprintln(os.Stderr, "atoi err", err)
		os.Exit(1)

	}

	result := ""
	for i := 0; i < n; i++ {
		text := readline(scanner)
		result = result + text
	}

	fmt.Println(result)
}

func readline(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "scan err")
		os.Exit(1)
	}
	return scanner.Text()
}
