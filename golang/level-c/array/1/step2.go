// 問題
// 以下の数列の全要素を改行区切りで出力してください。
// 5 1 3 4 5 12 6 8 1 3

package main

import "fmt"

func main() {
	arr := []int{5, 1, 3, 4, 5, 12, 6, 8, 1, 3}
	for _, v := range arr {
		fmt.Println(v)
	}
}
