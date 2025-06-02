// 問題
// 長さがわからない数列 a が入力されます。
// -1 が入力されるまで、受け取った数を改行区切りで出力してください。

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
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "failed to scan input")
		os.Exit(1)
	}
	inputs := strings.Fields(scanner.Text())

	// -1 が出るまでループし出力
	for _, input := range inputs {
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to atoi: %v \n", err)
			os.Exit(1)
		}
		fmt.Println(num)
		if num == -1 {
			return
		}
	}

}
