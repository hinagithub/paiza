// 問題
// 配列 A とその要素数 N と削除する要素の番号 n が与えられるので、
// A から A_n を削除し、A を要素数 N - 1 の配列とした後の A の各要素を出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	length := atoi(readline(scanner))
	inputs := make([]int, length)
	for i := 0; i < length; i++ {
		inputs[i] = atoi(readline(scanner))
	}

	// 対象要素を削除して出力
	// 例: deleteIndex = 3 のとき、3番目の要素（index 2）を削除する
	deleteIndex := atoi(readline(scanner))
	results := append(inputs[:deleteIndex-1], inputs[deleteIndex:]...)
	for _, r := range results {
		fmt.Println(r)
	}

}

func readline(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "failed to read input")
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
