// 問題
// 1 行目に整数 N と整数 K が与えられます。
// 2 行目に N 個の整数 a_i (1 ≤ i ≤ N) が半角スペース区切りで与えられます。
// K 番目の整数 a_K を出力してください。

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
	index, _ := strconv.Atoi(inputs[1])
	numbers := strings.Fields(readline(scanner))

	if index < 1 || index > len(numbers) {
		fmt.Fprintln(os.Stderr, "1から10の範囲で入力してください")
		os.Exit(1)
	}

	fmt.Println(numbers[index-1])

}

func readline(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Println("scan err")
		os.Exit(1)
	}
	return scanner.Text()
}
