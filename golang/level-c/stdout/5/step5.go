// 問題
// 10 個の数値が改行区切りで与えられます。
// これらの数値を半角スペース 2 つとバーティカルライン | 区切りで出力してください。
// ただし、末尾には改行を出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	sc := bufio.NewScanner(os.Stdin)
	inputs := [] string{}
	for i := 0; i < 10; i ++ {
		sc.Scan()
		inputs = append(inputs, sc.Text())
	}
	output := strings.Join(inputs, " | ")
	fmt.Println(output)
}