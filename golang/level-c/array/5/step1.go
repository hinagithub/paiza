// 問題
// 以下のような数列があります。

// 1 3 5 1 2 3 6 6 5 1 4

// この数列から数の重複をなくし、昇順にし改行区切りで出力してください。
// 数列を配列に格納し、並び替える操作や重複を削除する操作を考えて解いてみましょう。

package main

import (
	"fmt"
	"sort"
)

func main() {

	numbers := []int{1, 3, 5, 1, 2, 3, 6, 6, 5, 1, 4}
	uniqueNumbers := uniquify(numbers)
	sort.Ints(uniqueNumbers)
	for _, num := range uniqueNumbers {
		fmt.Println(num)
	}

}

func uniquify(numbers []int) []int {
	seen := make(map[int]bool)
	uniqueNumbers := []int{}
	for _, num := range numbers {
		if !seen[num] {
			seen[num] = true
			uniqueNumbers = append(uniqueNumbers, num)
		}
	}
	return uniqueNumbers
}
