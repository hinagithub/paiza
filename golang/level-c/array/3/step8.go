// 問題
// 以下のような配列があります。
// 1 2 2 1 2 1 2 1 1 1
// この中に含まれる 1 の個数を出力してください。

package main

import "fmt"

func main() {
	const target = 1
	numbers := []int{1, 2, 2, 1, 2, 1, 2, 1, 1, 1}

	count := 0
	for _, num := range numbers {
		if num == target {
			count++
		}
	}
	fmt.Println(count)

}
