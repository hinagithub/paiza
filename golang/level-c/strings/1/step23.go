// 数字からなる文字列 S が与えられるので、 S の先頭（左側）から S を見て行ったときにその数字が現れたのが 2 回目以降の場合、その数字を削除した文字列（重複した数字を削除した文字列）を求めてください。

// 例
// ・ S = "1234567890" のとき
// 数字に重複はないので "1234567890" が答えとなります。

// ・ S = "112123123412345" のとき
// "12345" が答えとなります。
// ・ 1 は 1 文字目に初めて現れるので、それ以降の 1 は読み飛ばします。
// ・ 2 は 3 文字目に初めて現れ、それ以降の 2 は読み飛ばします。
// ・ 3 は 6 文字目に初めて現れ、それ以降の 3 は読み飛ばします。
// ・ 4 は 10 文字目に初めて現れ、それ以降の 4 は読み飛ばします。
// ・ 5 は 15 文字目に初めて現れ、それ以降の 5 は読み飛ばします。

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

	inputs := []rune(scanner.Text())
	seen := make(map[rune]bool)
	result := ""

	for _, input := range inputs {
		if !seen[input] {
			seen[input] = true
			result = result + string(input)
		}
	}
	fmt.Println(result)
}
