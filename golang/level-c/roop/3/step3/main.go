// 問題
// 整数 N が与えられます。
// 九九の N の段を半角スペース区切りで出力してください。

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	tableN := make([]string, 9)
	for i := 1; i <= 9; i++ {
		tableN[i-1] = strconv.Itoa(n * i)
	}
	fmt.Println(strings.Join(tableN, " "))
}
