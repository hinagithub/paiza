// 問題
// 1 行目に整数 N, M が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// 左から M 番目の要素を削除し、削除した後の N - 1 個の要素を改行区切りで出力してください。
// なお、左端を 1 番目とします。

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
	inputs := strings.Fields(readline(scanner))
	targetIndex, err := strconv.Atoi(inputs[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	strNumbers := strings.Fields(readline(scanner))

	results := []string{}
	results = append(results, strNumbers[:targetIndex-1]...)
	results = append(results, strNumbers[targetIndex:]...)

	for _, result := range results {
		fmt.Println(result)
	}
}

func readline(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Println("scan errer")
		os.Exit(1)
	}
	return scanner.Text()
}
