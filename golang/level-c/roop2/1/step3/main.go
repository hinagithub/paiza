// 問題
// N 個の整数 a_1, a_2, ..., a_N が与えられます。
// この N 個の整数のうち、a_1 から順に奇数か偶数か判定し、奇数の場合のみ改行区切りで出力してください。
// また、N 個の整数には奇数が少なくとも 1 つ含まれています。

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
	length := atoi(readline(scanner))
	inputs := strings.Fields(readline(scanner))
	for i := 0; i < length; i++ {
		num := atoi(inputs[i])

		// 奇数の場合のみ出力
		if num%2 == 1 {
			fmt.Println(num)
		}
	}
}

func readline(scanner *bufio.Scanner) string {
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
