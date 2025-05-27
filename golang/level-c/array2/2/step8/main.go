// 問題
// 配列 A の要素数 N と配列 A の各要素 A_1, A_2, ...,
// A_N が与えられるので、A の要素の全てのペアについてのかけ算の結果を出力してください。

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

	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			fmt.Println(inputs[i] * inputs[j])
		}
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
