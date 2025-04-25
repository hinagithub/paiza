// 問題
// 偶数 N が入力されます。
// まず、 1 行目には 1 以上 N / 2 以下の数値を半角スペース区切りですべて出力してください。
// 次に、 2 行目には N / 2 + 1 以上 N 以下の数値を半角スペース区切りですべて出力してください。
// 各行の末尾には、半角スペースの代わりに改行を入れてください。

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
	limit, _ := strconv.Atoi(sc.Text())
	border := limit / 2
	for i := 1; i <= border; i++ {
		fmt.Print(i)
	}
	fmt.Println()
	for i := border; i <= limit; i++ {
		fmt.Print(i)
	}
}
