// 問題
// 整数 K が与えられます。
// 数列 a が
// 1 3 5 4 6 2 1 7 1 5
// で与えられるとき、数列 a の K 番目の要素を出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers := []int{1, 3, 5, 4, 6, 2, 1, 7, 1, 5}
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		k, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, "入力が整数ではありません:", err)
			os.Exit(1)
		}

		if k < 1 || k > len(numbers) {
			fmt.Fprintln(os.Stderr, "1から10の範囲で入力してください")
			os.Exit(1)
		}

		fmt.Println(numbers[k-1])
	}
}
