// 問題
// 1 行目に整数 N, K が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// N 個の整数のうち、K 以上の数をすべて、入力された順に改行区切りで出力してください。
// また、K 以上の数が一個もない場合は、何も出力しなくても問題ありません。

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
	scanner.Scan()
	inputs := strings.Fields(scanner.Text())
	border, err := strconv.Atoi(inputs[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "atoi err: %d", err)
		os.Exit(1)
	}
	scanner.Scan()

	for _, strNumber := range strings.Fields(scanner.Text()) {
		num, err := strconv.Atoi(strNumber)
		if err != nil {
			fmt.Fprintf(os.Stderr, "atoi err: %d", err)
			os.Exit(1)
		}
		if num >= border {
			fmt.Println(num)
		}
	}
}
