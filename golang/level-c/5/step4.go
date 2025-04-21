// 問題
// 10 個の数値が半角スペース区切りで与えられます。
// これらの数値をカンマ区切りで出力してください。
// ただし、末尾にはカンマではなく改行を出力してください。

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
	output := strings.Join(inputs, ",")
	fmt.Println(output)
}