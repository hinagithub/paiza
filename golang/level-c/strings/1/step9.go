// 問題
// 文字列 S , T と、整数 N が与えられるので、 S の N 文字目の後ろに T を挿入した文字列を出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 挿入前のテキスト
	scan(scanner)
	text := scanner.Text()

	// 挿入するテキスト
	scan(scanner)
	addition := scanner.Text()

	// 挿入する場所のインデックス番号
	scan(scanner)
	index, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Fprintln(os.Stderr, "atoi err")
		os.Exit(1)
	}

	// 文字列を rune スライスに変換（マルチバイト文字対応）
	runes := []rune(text)
	if index < 0 || index > len(runes) {
		fmt.Fprintln(os.Stderr, "index out of range")
		os.Exit(1)
	}

	// 結果出力
	result := string(runes[:index]) + addition + string(runes[index:])
	fmt.Println(result)
}

func scan(scanner *bufio.Scanner) {
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "scan err")
		os.Exit(1)
	}
}
