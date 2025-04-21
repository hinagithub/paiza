// 問題
// 3 つの文字列が改行区切りで与えられます。
// これらの文字列をバーティカルライン | 区切りで出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	sc := bufio.NewScanner(os.Stdin)
	inputs := []string{}
	for i := 0; i < 3; i++ {
		sc.Scan()
		inputs = append(inputs, sc.Text())
	}
	result := strings.Join(inputs,"|")
	fmt.Println(result)
}