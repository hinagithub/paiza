// 問題
// N 個の整数 K_1, K_2, ..., K_N が与えられます。
// これらを受け取り、改行区切りで出力してください。

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
	readline(scanner) // 1行目の桁数は使わないのでスキップ
	inputs := strings.Fields(readline(scanner))
	for _, input := range inputs {
		num := atoi(input)
		fmt.Println(num)
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
