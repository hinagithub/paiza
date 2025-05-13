// 問題
// 正しい数式を表す文字列 S が与えられるので、その数式を計算した結果を出力してください。
// ただし、出てくる計算は足し算・引き算のみとし、数式に項として出てくる数字は全て 1 桁の正の数であるものとします。
// 例として、1-2, 2+1-7 のような文字列は与えられる可能性がありますが、1+(-2) や -4-2+4, 5+-2 のような文字列は与えられないことが保証されている点に注意してください。

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
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("scan err")
		os.Exit(1)
	}
	text := scanner.Text()

	result := 0
	prevSign := '+'

	for _, char := range text {
		switch char {
		case '+', '-':
			prevSign = char
		default:
			num := atoi(char)
			if prevSign == '+' {
				result += num
			} else {
				result -= num
			}
		}
	}

	fmt.Println(result)
}

func atoi(char rune) int {
	num, err := strconv.Atoi(string(char))
	if err != nil {
		fmt.Fprintf(os.Stderr, "atoi error: %v\n", err)
		os.Exit(1)
	}
	return num
}
