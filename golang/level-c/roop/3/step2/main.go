// 問題
// 九九の 8 の段を半角スペース区切りで出力してください。

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	table8 := make([]string, 9)
	for i := 1; i <= 9; i++ {
		table8[i-1] = strconv.Itoa(8 * i)
	}
	fmt.Println(strings.Join(table8, " "))
}
