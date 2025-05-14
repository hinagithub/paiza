// 問題
// 整数 N が与えられます。
// 整数 N が、以下の配列の左から何番目にあるか出力してください。
// 左端を 1 番目とし、N は以下の配列に必ず含まれるものとします。

// 1 5 9 7 3 2 4 8 6 10

package main

import "fmt"

func main() {
	numbers := []int{1, 5, 9, 7, 3, 2, 4, 8, 6, 10}
	var target int
	fmt.Scan(&target)

	for i, num := range numbers {
		if num == target {
			fmt.Println(i + 1)
		}
	}
}
