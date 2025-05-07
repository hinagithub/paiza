// 問題
// 文字列 S と整数 i と文字 c が与えられるので、S の i 文字目を c に書き換えたものを出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// テキストを取得
	scanner.Scan()
	text := scanner.Text()

	// インデックスを取得
	scanner.Scan()
	inputs := strings.Fields(scanner.Text())
	index, err := strconv.Atoi(inputs[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "strconv err")
		os.Exit(1)
	}
	// 文字の位置からインデックス番号に-1する
	index--
	if index < 0 || index > len(text) {
		fmt.Println("index length out of range")
		os.Exit(1)
	}

	// 置き換え文字列を取得
	replacement := []rune(inputs[1])
	if len(replacement) != 1 {
		fmt.Println("replacement word must be 1 char")
		os.Exit(1)
	}

	// 置換して出力
	runes := []rune(text)
	runes[index] = replacement[0]
	fmt.Println(string(runes))
}
