// 問題
// 1 行目に整数 N が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// N 個の整数のうち、5 以上の数をすべて、入力された順に改行区切りで出力してください。
// また、5 以上の数が一個もない場合は、何も出力しなくても問題ありません。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 基準値
	const border = 5
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan() // 1行目は使わないのでスキップ
	scanner.Scan()
	inputs := strings.Fields(scanner.Text())

	for _, input := range inputs {
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "atoi err: %d", err)
			os.Exit(1)
		}

		// 基準値以上の数だけ出力
		if num >= border {
			fmt.Println(num)
		}
	}

}
