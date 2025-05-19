// 問題
// 1 行目に整数 N, M が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// N 個の整数の末尾に整数 M を挿入し、改行区切りで出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputs := strings.Fields(scanner.Text())
	newWord := inputs[1]

	scanner.Scan()
	strNumbers := strings.Fields(scanner.Text())

	strNumbers = append(strNumbers, newWord)
	for _, srtNumber := range strNumbers {
		fmt.Println(srtNumber)
	}
}
