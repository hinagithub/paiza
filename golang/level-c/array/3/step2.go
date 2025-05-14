// 問題
// 整数 N が与えられます。
// 以下の配列に、整数 N が含まれているなら Yes、含まれていないなら No を出力してください。
// 5 12 6 84 14 25 44 3 7 20

package main

import "fmt"

func main() {
	const (
		Yes = "Yes"
		No  = "No"
	)
	var target int
	fmt.Scan(&target)

	numbers := []int{5, 12, 6, 84, 14, 25, 44, 3, 7, 20}
	found := false
	for _, num := range numbers {
		if num == target {
			found = true
			break
		}
	}

	if found {
		fmt.Println(Yes)
	} else {
		fmt.Println(No)
	}
}
