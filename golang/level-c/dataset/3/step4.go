// 問題
// paiza 商店では N 個の商品が売られており、i 番目の商品の名前は A_i で、価格は B_i です。
// あなたは M 個の商品名が書かれたお買い物リスト S を持っています。
// リストに書かれているそれぞれの商品について、paiza 商店での価格を出力してください。
// リストには paiza 商店が扱っていない商品も書かれている可能性がありますが、
// その場合は価格の代わりに -1 を出力してください。

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

	// item数を取得
	scanner.Scan()
	inputs := strings.Fields(scanner.Text())
	itemCount := atoi(inputs[0])

	// 販売しているアイテムリストを取得
	shopItems := make(map[string]int, itemCount)
	for i := 0; i < itemCount; i++ {
		scanner.Scan()
		items := strings.Fields(scanner.Text())
		name := items[0]
		price := atoi(items[1])
		shopItems[name] = price
	}

	// ほしいものと一致する商品の価格を出力
	for scanner.Scan() {
		want := scanner.Text()
		if price, exists := shopItems[want]; exists {
			fmt.Println(price)
		} else {
			fmt.Println(-1)
		}
	}
}

func atoi(text string) int {
	number, err := strconv.Atoi(text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to convert to int: %s\n", text)
		os.Exit(1)
	}
	return number
}
