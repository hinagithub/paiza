// 問題
// 配列 A と追加する位置 n と追加する要素 B が与えられるので、B を A_n の後ろに追加した後の A を出力してください。

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
	// 要素数を取得
	length := atoi(readline(scanner))

	// 配列を取得
	inputs := make([]string, length+1)
	for i := 0; i < length; i++ {
		inputs[i] = readline(scanner)
	}

	// 追加する要素の要素番号と値を取得
	addingInfos := strings.Fields(readline())
	addingIndex := aoit(addingInfos[0])
	addingValue := addingInfos[1]

	// 要素を追加して出力
	results := append(inputs[:addingIndex], addingValue)
	results := append(results, inputs[addingIndex+1:])
	for _, result := range results {
		fmt.Println(result)
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
		fmt.Fprintf(os.Stderr, "failed to atoi: %v, \n")
		os.Exit(1)
	}

	return num
}
