// 問題
// paiza 研究所では実験データをわかりやすい形式で保存したいのですが、
// 機械によって得られる数値には、次のような表記ミスがあることがわかっています。

// ・ ミス 1
// 先頭に必要のない 0 がいくつかついてしまう
// ・ 本来 1 である数値が 0001 と表記されてしまう
// ・ 本来 0.001 である数値が 00.001 と表記されてしまう

// ・ ミス 2
// 小数である数値の末尾に必要のない 0 がいくつかついてしまう
// ・ 本来 0.1 である数値が 0.10 と表記されてしまう

// ・ ミス 3
// 小数である数値に小数点が複数個ついてしまう
// ただし、 1 つめの小数点の位置が正しい小数点の位置であるものとします
// ・ 本来 0.123 である数値が 0.1.2.3 と表記されてしまう。

// 表記が正しくない可能性のある数値を表す文字列 S が与えられるので、その数値を正しい表記にしてください。
// S を数値として扱うと上手く処理が行えないので気をつけてください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("scan err")
		os.Exit(1)
	}
	text := scanner.Text()

	fixed := ""

	if firstDotIndex(text) == -1 {
		// 整数の場合
		fixed = suppress(text)
	} else {

		// 少数の場合
		fixed = miss3Fix(text)
		fixed = miss1Fix(fixed)
		fixed = miss2Fix(fixed)
	}

	fmt.Println(fixed)
}

// 最初のドットのインデックスを取得
func firstDotIndex(text string) int {
	for i, char := range text {
		if char == '.' {
			return i
		}
	}
	return -1
}

// 整数のゼロサプレス
func suppress(text string) string {

	result := ""
	seenFirstNum := false

	for _, char := range text {
		if seenFirstNum {
			result = result + string(char)
			continue
		}

		if char != '0' {
			seenFirstNum = true
			result = result + string(char)
			continue
		}
	}
	return result
}

// ミス3の修正(小数点が複数あるミス)
func miss3Fix(text string) string {
	dotIndex := firstDotIndex(text)
	if dotIndex == -1 {
		return text
	}
	result := ""
	for i, char := range text {
		if char == '.' && i != dotIndex {
			continue
		}
		result = result + string(char)
	}
	return result
}

// ミス1修正(整数部の0サプレス)
func miss1Fix(text string) string {
	dotIndex := firstDotIndex(text)
	result := ""
	seenFirstNum := false

	for i, char := range text {

		// 少数部なら比較しない
		if i > dotIndex {
			result = result + string(char)
			continue
		}

		// ドット直前の0は許容
		if dotIndex-1 == i {
			result = result + string(char)
			continue
		}

		if seenFirstNum {
			result = result + string(char)
			continue
		}

		if char != '0' {
			seenFirstNum = true
			result = result + string(char)
			continue
		}
	}
	return result
}

// ミス2の修正(後方0の重複排除)
func miss2Fix(text string) string {
	dotIndex := firstDotIndex(text)
	end := len(text)
	for i := len(text) - 1; i > dotIndex; i-- {
		if text[i] != '0' {
			end = i + 1
			break
		}
	}
	return text[:end]
}
