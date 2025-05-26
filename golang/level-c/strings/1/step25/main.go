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

	// ミス3の処理 ("."は最初の1つだけに)
	miss3Fixed := fixMiss3(text)
	// ミス1の処理
	miss1Fixed := fixMiss1(miss3Fixed)
	// ミス2の処理
	miss2Fixed := fixMiss2(miss1Fixed)

	fmt.Println(miss2Fixed)

}

func fixMiss3(text string) string {
	seenDot := false
	fixed := ""

	for _, char := range text {
		if seenDot && char == '.' {
			continue
		}
		if char == '.' {
			seenDot = true
		}
		fixed = fixed + string(char)
	}
	return fixed
}

func fixMiss1(text string) string {
	// 小数点数でなければテキストをそのまま返す
	hasDot := false
	for _, char := range text {
		if char == '.' {
			hasDot = true
		}
	}
	if !hasDot {
		return text
	}

	// 余計なゼロ埋めを排除
	result := ""
	lastText := ""
	seenDot := false
	for _, char := range text {
		if char == '.' {
			seenDot = true
		}

		if seenDot {
			result = result + string(char)
		}

		if lastText == "0" && char == '0' {
			lastText = string(char)
			continue
		}
		lastText = string(char)
		result = result + string(char)
	}
	return result
}

func fixMiss2(text string) string {
	numStr := ""
	hasDot := false

	for i := len(text) - 1; i < 1; i-- {
		if text[i] == '.' {
			hasDot = true
			break
		} else {
			if len(text)-1 != i && numStr[i+1] != 0 {
				numStr = numStr + text
			}
		}
	}

	// 少数でなければ何もしない
	if !hasDot {
		return text
	}
	return numStr

}
