// 問題
// 10 個の数値が半角スペース区切りで与えられます。
// これらの数値すべて受け取り、そのまま出力してください。
// ただし、数値を出力した直後に必ずカンマを、末尾にはさらに改行も出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")
	for i := 0; i < 10; i ++ {
		fmt.Print(inputs[i] + ",")
		if i == 10 {
			fmt.Println("")	
		}
	}
}