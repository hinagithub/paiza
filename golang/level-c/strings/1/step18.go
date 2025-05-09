// 問題
// 文字列 S が与えられるので、 S が回文かどうかを判定してください。
// なお、回文とは、前から読んでも後ろから読んでも同じ文字列になるような文字列のことをいいます。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const (
		Yes = "YES"
		No  = "NO"
	)

	// テキストの取得
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "scan err")
		os.Exit(1)
	}
	text := scanner.Text()
	runes := []rune(text)
	length := len(runes)
	isPalindrome := Yes

	for i := 0; i < length/2; i++ {
		if runes[i] != runes[length-1-i] {
			isPalindrome = No
			break
		}
	}
	fmt.Println(isPalindrome)
}
