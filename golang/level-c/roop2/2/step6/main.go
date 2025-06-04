// 問題
// 10 進数で表された整数 N, M が与えられます。
// N を M 進数に変換して出力してください。

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(strconv.FormatInt(int64(n), m))
}
