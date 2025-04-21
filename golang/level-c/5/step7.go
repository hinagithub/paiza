// 問題
// 大きな数値Nが入力されます。
// 位の小さい方から 3 けたごとにカンマ区切りで出力してください。
// ただし、Nのけた数は 3 の倍数とは限りません。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	sc:= bufio.NewScanner(os.Stdin)
	sc.Scan()
	n := sc.Text()
	charactors := strings.Split(n, "")
	results := []string{}

	for i, char := range charactors {
		if i != 0 && ( len(charactors)-i )  % 3  == 0 {
			results = append(results,",")
		}
		results = append(results, char)
	}
	result := strings.Join(results,"")
	fmt.Println(result)
}