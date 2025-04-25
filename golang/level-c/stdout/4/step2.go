// 問題
// 数値 N が入力されます。
// 1 から N までの数値をすべて、改行区切りで出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n,_ := strconv.Atoi(sc.Text())
	for i := 1; i <= n ; i ++ {
		fmt.Println(i)
	}
}