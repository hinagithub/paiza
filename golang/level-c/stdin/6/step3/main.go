// 問題
// 1 行目で、整数 N と、続けて N 個の整数 a_1, ... , a_N が半角スペース区切りで与えられます。
// a_1, ... , a_N を改行区切りで出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fields := strings.Fields(sc.Text())
	// 1つ目はN（整数の数）なのでスキップ
	for _, numStr := range fields[1:] {
		fmt.Println(numStr)
	}
}
