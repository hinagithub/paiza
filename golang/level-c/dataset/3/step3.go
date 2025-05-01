// 問題
// 文字列が N 個与えられます。各文字列の出現回数を文字列の辞書順に出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// 標準入力から文字列の数を読み込む
	var n int
	fmt.Scan(&n)

	wordCount := make(map[string]int)

	// 入力された文字列を読み込み、カウント
	for i := 0; i < n; i++ {
		scanner.Scan()
		word := scanner.Text()
		wordCount[word]++
	}

	// 辞書順に並べるためキー（単語）をスライスに取り出す
	words := make([]string, 0, len(wordCount))
	for word := range wordCount {
		words = append(words, word)
	}
	sort.Strings(words)

	// 辞書順に出力
	for _, word := range words {
		fmt.Printf("%s %d\n", word, wordCount[word])
	}

}
