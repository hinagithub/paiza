// 問題
// 大きな数値 N が入力されます。
// 3 けたごとにカンマ区切りで出力してください。
// ただし、N のけた数は 3 の倍数です。

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