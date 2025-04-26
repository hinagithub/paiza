// 問題
// 1 行目で整数 N が与えられます。
// 2 行目以降で、N 個の整数 a_1, ... , a_N が N 行で与えられます。
// a_1, ... , a_N を改行区切りで出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	len, _ := strconv.Atoi(sc.Text())
	for i := 0; i < len; i++ {
		sc.Scan()
		fmt.Println(sc.Text())
	}
}
