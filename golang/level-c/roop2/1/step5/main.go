// 問題
// N 個の整数 a_1, a_2, ..., a_N が与えられます。
// この N 個の整数について、a_1 から順に 奇数か偶数か判定し、奇数なら odd 、偶数なら even を改行区切りで出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const (
		Odd  = "odd"  // 奇数
		Even = "even" // 偶数
	)
	scanner := bufio.NewScanner(os.Stdin)
	readLine(scanner) // 1行目の要素数は使用しないのでスキップ
	inputs := strings.Fields(readLine(scanner))

	// 奇数・偶数を出力
	for _, input := range inputs {
		num := atoi(input)
		if num%2 == 0 {
			fmt.Println(Even)
		} else {
			fmt.Println(Odd)
		}
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
