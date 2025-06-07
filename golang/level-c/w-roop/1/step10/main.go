// 整数 N が与えられるので、次の規則に従って N 行の出力をしてください。
//  N 行のうち、 i 行目では、1 から i までの数字を半角スペース区切りで出力してください。
// 例として、 N = 3 のとき、出力は次の通りになります。

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		row := []string{}
		for j := 1; j <= i; j++ {
			row = append(row, strconv.Itoa(j))
		}
		fmt.Println(strings.Join(row, " "))
	}
}
