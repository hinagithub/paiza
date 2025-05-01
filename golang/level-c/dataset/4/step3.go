// 問題
// N 個の要素からなる数列 A が与えられます。
// 2 ≦ i ≦ N の各 i に対して、A_i と同じ値が A_1 から A_{i-1} の間にあるかどうかを判定してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const (
		Yes = "Yes"
		No  = "No"
	)
	scanner := bufio.NewScanner(os.Stdin)

	// 1行目の要素数は使わないのでスキップ
	scan(scanner)

	// 入力値を取得
	scan(scanner)
	texts := strings.Fields(scanner.Text())
	founds := make(map[string]struct{})

	// 最初の要素は記録だけ
	founds[texts[0]] = struct{}{}

	for _, text := range texts[1:] {
		_, exists := founds[text]

		if exists {
			fmt.Println(Yes)
		} else {
			founds[text] = struct{}{}
			fmt.Println(No)
		}

	}
}
func scan(scanner *bufio.Scanner) {
	if !scanner.Scan() {
		fmt.Println(os.Stderr, "scan err")
		os.Exit(1)
	}
}
