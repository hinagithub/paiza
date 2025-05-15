// 問題
// 1 行目に整数 N, M が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// 3 行目に M 個の整数 b_1, b_2, ..., b_M が与えられます。
// N 個の整数 a_1, a_2, ..., a_N の後ろに M 個の整数 b_1, b_2, ..., b_M を連結させ、
// 改行区切りで出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	readline(scanner) // 1行目は使用しないので飛ばす
	arr1 := strings.Fields(readline(scanner))
	arr2 := strings.Fields(readline(scanner))

	results := append(arr1, arr2...)
	for _, str := range results {
		fmt.Println(str)
	}

}

func readline(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Println("scan error")
		os.Exit(1)
	}
	return scanner.Text()
}
