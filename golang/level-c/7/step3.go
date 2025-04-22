// 問題
// 自然数 N が与えられます。1 ≦ i ≦ N の各 i について、i 行目には以下の数列を出力してください。
// * 1 以上 i 以下の数値をすべて、半角スペース区切りで出力してください

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	for i := 1; i <= n; i++ {
		results := []string{}
		for j := 1; j <= i; j++ {
			results = append(results, strconv.Itoa(j))
		}
		fmt.Println(strings.Join(results, " "))
	}
}
