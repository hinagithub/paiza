// 3 つの文字列

// "zaipa" izapa paiza

// があります。辞書順に並べ替え、改行区切りで出力してください。
// この三つの文字列を配列に格納し、その配列を並び替える操作を考えて解いてみましょう。

package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"zaipa", "izapa", "paiza"}
	sort.Strings(words)
	for _, word := range words {
		fmt.Println(word)
	}
}

// func main() {

// 	words := []string{"zaipa", "izapa", "paiza"}
// 	for i := len(words) - 1; i >= 0; i-- {
// 		maxIndex, maxValue := findMax(words[:i+1])
// 		words[maxIndex] = words[i]
// 		words[i] = maxValue
// 	}
// 	for _, word := range words {
// 		fmt.Println(word)
// 	}

// }

// func findMax(words []string) (int, string) {
// 	maxValue := words[0]
// 	maxIndex := 0
// 	for i, word := range words {
// 		if maxValue <= word {
// 			maxValue = word
// 			maxIndex = i
// 		}
// 	}
// 	return maxIndex, maxValue
// }
