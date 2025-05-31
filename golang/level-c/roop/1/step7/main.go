// 問題
// a ~ z のアルファベットを、改行区切りで出力してください。

package main

import "fmt"

func main() {
	alphabets := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}

	for _, alphabet := range alphabets {
		fmt.Println(alphabet)
	}

	// こう書ける
	// 	for ch := 'a'; ch <= 'z'; ch++ {
	// 		fmt.Printf("%c\n", ch)
	// 	}
}
