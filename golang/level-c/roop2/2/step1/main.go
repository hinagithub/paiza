// 問題
// 10 進数で表された整数 N が与えられます。
// 整数 N の各桁の和を計算し、出力してください。

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var strN string
	fmt.Scan(&strN)

	sum := 0
	for _, rune := range strN {
		num, err := strconv.Atoi(string(rune))
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to atoi: %v \n", err)
			os.Exit(1)
		}
		sum += num
	}
	fmt.Println(sum)

}
