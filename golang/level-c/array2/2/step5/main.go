// 問題
// 配列 A と追加する位置 n と追加する要素 B が与えられるので、
// B を A_n の後ろに追加した後の A を出力してください。

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

	// 配列を取得
	length := atoi(readline(scanner))
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		numbers[i] = atoi(readline(scanner))
	}
	// 追加する要素の情報を取得
	addingInfo := strings.Fields(readline(scanner))
	index := atoi(addingInfo[0])
	val := atoi(addingInfo[1])

	// 要素を追加し出力
	results := append(numbers[:index], append([]int{val}, numbers[index:]...)...)
	for _, r := range results {
		fmt.Println(r)
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
		fmt.Fprintf(os.Stderr, "failed to atoi : %v \n", err)
		os.Exit(1)
	}
	return num
}
