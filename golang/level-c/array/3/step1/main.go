// 問題
// 以下のような配列があります。

// 10 13 21 1 6 51 10 8 15 6

// この中に、6 が含まれているなら Yes、含まれていないなら No を出力してください。

package main

import "fmt"

func main() {
	const (
		Target = 6
		Yes    = "Yes"
		No     = "No"
	)
	numbers := []int{10, 13, 21, 1, 6, 51, 10, 8, 15, 6}

	found := false
	for _, number := range numbers {
		if number == Target {
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
