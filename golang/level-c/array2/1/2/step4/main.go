// 問題
// 配列 A と追加する要素 B が与えられるので、B を A の末尾に追加したのち、A の全ての要素を出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 要素数を取得
	length, err := strconv.Atoi(readline(scanner))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to atoi: %v \n", err)
		os.Exit(1)
	}

	// 配列を取得
	inputs := make([]string, length)
	for i := 0; i < length; i++ {
		inputs[i] = readline(scanner)
	}

	// 配列末尾に追加して出力
	additionalWord := readline(scanner)
	inputs = append(inputs, additionalWord)
	for _, str := range inputs {
		fmt.Println(str)
	}

}

func readline(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "failed to scan input")
		os.Exit(1)
	}
	return scanner.Text()
}
