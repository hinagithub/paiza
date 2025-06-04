// 問題
// 10 進数で表された整数 N が与えられます。
// N を 2 進数に変換して出力してください。

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	binary := strconv.FormatInt(int64(n), 2)
	fmt.Println(binary)
}
