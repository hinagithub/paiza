// 問題
// N 個の要素からなる数列 A と、整数 B が与えられます。B が A に含まれているかどうかを判定してください。

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
		Yes = "Yes"
		No  = "No"
	)

	// 対象の値を取得
	scanner := bufio.NewScanner(os.Stdin)
	text := readline(scanner)
	inputs := strings.Fields(text)
	target := atoi(inputs[1])

	// 探索元の配列を取得
	numberText := readline(scanner)
	numberTexts := strings.Fields(numberText)

	// 判定
	seen := No
	for _, text := range numberTexts {
		number := atoi(text)
		if number == target {
			seen = Yes
			break
		}
	}
	fmt.Println(seen)
}

func readline(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Println(os.Stderr, "scan err")
		os.Exit(1)
	}
	text := scanner.Text()
	return text
}

func atoi(text string) int {
	number, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("atoi err")
		os.Exit(1)
	}
	return number
}
