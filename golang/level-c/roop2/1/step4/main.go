// 問題
// N 個の整数 a_1, a_2, ..., a_N が与えられます。
// この N 個の整数のうち、a_1 から順に 3 で割り切れるか判定し、割り切れる場合のみ改行区切りで出力してください。
// また、N 個の整数には 3 で割り切れる数が少なくとも 1 つ含まれています。

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

	// 3で割り切れる場合のみ出力
	for _, input := range inputs {
		num := atoi(input)
		if num%3 == 0 {
			fmt.Println(input)
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
