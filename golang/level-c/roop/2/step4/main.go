// 問題
// 長さ N の数列 a (a_1, a_2, ..., a_N) と b (b_1, b_2, ..., b_N) が与えられます。
// a の各要素から b の各要素を引いた結果 (a_1 - b_1, a_2 - b_2, ..., a_N - b_N) を、改行区切りで出力してください。
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
	readline(scanner) // 1行目の要素数Nは使用しないのでスキップ
	arrA := strings.Fields(readline(scanner))
	arrB := strings.Fields(readline(scanner))

	// 配列長さチェック
	if len(arrA) != len(arrB) {
		fmt.Fprintln(os.Stderr, "input arrA and arrB length must be same")
		os.Exit(1)
	}

	for i := 0; i < len(arrA); i++ {
		numA := atoi(arrA[i])
		numB := atoi(arrB[i])
		fmt.Println(numA - numB)
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
		fmt.Fprintf(os.Stderr, "failed to atoi: %v \n", err)
		os.Exit(1)
	}
	return num
}
