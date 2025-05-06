// 問題
// 文字列 S と整数 i が与えられるので、 S の i 文字目を出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 自分の得意な言語で
	// Let's チャレンジ！！
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println(os.Stderr, "scan err")
		os.Exit(1)
	}
	text := scanner.Text()
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "scan err")
		os.Exit(1)
	}
	i, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Fprintln(os.Stderr, "atoi err", err)
		os.Exit(1)
	}

	runes := []rune(text)
	if i < 1 || i > len(runes) {
		fmt.Println("index err")
		os.Exit(1)
	}

	fmt.Println(string(runes[i-1]))
}
