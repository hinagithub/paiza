// 問題
// 整数 N が与えられます。
// 以下の配列に含まれる N の個数を出力してください。
// また、N は以下の配列に 1 個以上含まれるものとします。

// 1 2 5 1 4 3 2 5 1 4

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 5, 1, 4, 3, 2, 5, 1, 4}
	var target int
	fmt.Scan(&target)

	count := 0
	for _, num := range numbers {
		if num == target {
			count++
		}
	}
	fmt.Println(count)
}
