// 問題
// 数値 N (N = 1 または 2) が入力されます。
// N = 1 の場合は 1 を、N = 2 の場合は 1 と 2 を改行区切りで出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main (){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n,_ := strconv.Atoi(sc.Text())
	for i :=1; i <=n; i ++ {
		fmt.Println(i)
	}
}
