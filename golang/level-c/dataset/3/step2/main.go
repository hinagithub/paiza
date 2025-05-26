// 問題
// 長さ N の文字列 S が与えられます。
// S に含まれている各文字の出現回数をそれぞれ求め、「a」の出現回数、「b」の出現回数、...、
//「z」の出現回数をこの順に半角スペース区切りで1行に出力してください。

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

	// 1行目の長さは使わないので読み飛ばす
	scanner.Scan()

	// 2行目：文字列を取得して小文字化
	scanner.Scan()
	input := strings.ToLower(scanner.Text())
	runes := []rune(input)

	alphabets := []rune("abcdefghijklmnopqrstuvwxyz")
	results := []int{}

	for _, alphabet := range alphabets {
		results = append(results, countRune(runes, alphabet))
	}

	fmt.Println(strings.Join(itoaAll(results), " "))
}

func countRune(runes []rune, target rune) int {
	count := 0
	for _, r := range runes {
		if r == target {
			count++
		}
	}
	return count
}

func itoaAll(numbers []int) []string {
	texts := []string{}
	for _, number := range numbers {
		texts = append(texts, strconv.Itoa(number))
	}
	return texts
}
