// 問題
// 以下の数列の全要素を改行区切りで出力してください。
// 8 1 3 3 8 1 1 3 8 8

package main

import "fmt"

func main() {
	arr := []int{8, 1, 3, 3, 8, 1, 1, 3, 8, 8}
	for _, v := range arr {
		fmt.Println(v)
	}
}
