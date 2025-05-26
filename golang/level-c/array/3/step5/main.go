//,問題
// 以下のような配列があります。
// 1 10 2 9 3 8 4 7 5 6
// この配列の中で、8 が左から何番目にあるか出力してください。
// 左端を 1 番目とします。

package main

import "fmt"

func main() {
	const Target = 8
	numbers := []int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}
	for i, num := range numbers {
		if num == Target {
			fmt.Println(i + 1)
		}
	}
}
