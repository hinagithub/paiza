// 問題
// 正しい数式を表す文字列 S が与えられるので、その数式を計算した結果を出力してください。ただし、出てくる計算は足し算・引き算のみとします。

// ・ 例
// ・ S = "1+1"
// 答えは 2 となります。

// ・ S = "4+3-2+1"
// 答えは 6 となります。

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
		fmt.Println("scan err")
		os.Exit(1)
	}
	text := scanner.Text()

	numStrs := strings.FieldsFunc(text, splitBySign)
	signs := []rune{}
	for _, char := range text {
		if char == '+' {
			signs = append(signs, '+')
		} else if char == '-' {
			signs = append(signs, '-')
		}
	}
	result := 0
	for i, sign := range signs {
		if i == 0 {
			num, _ := strconv.Atoi(numStrs[i])
			result += num
		}
		num, _ := strconv.Atoi(numStrs[i+1])
		if sign == '+' {
			result += num
		} else {
			result -= num
		}
	}
	fmt.Println(result)
}

func splitBySign(r rune) bool {
	return r == '+' || r == '-'
}
