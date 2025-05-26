// N 個の要素からなる数列 A が与えられます。数列 A に対し、次の 3 つの操作を行うプログラムを作成してください。

// ・ push_back x : A の末尾に x を追加する
// ・ pop_back : A の末尾を削除する
// ・ print : A を半角スペース区切りで1行に出力する

// 例えば、入力例 1 において、数列 A は最初「1 2 3」です。最初の操作は「push_back 10」なので、末尾に 10 を追加して「1 2 3 10」となります。 2 つ目の操作は「push_back 12」なので、「1 2 3 10 12」となります。 3 つ目の操作は「print」なので「1 2 3 10 12」を出力します。 4 つ目の操作は「pop_back」なので末尾の「12」を削除して、「1 2 3 10」となります。最後の操作は「print」なので「1 2 3 10」を出力します。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	// 1行目は読み飛ばす
	sc.Scan()

	// ベース配列を取得
	sc.Scan()
	tokens := strings.Fields(sc.Text())

	for sc.Scan() {
		inputs := strings.Fields(sc.Text())
		command := inputs[0]
		switch command {
		case "0": // 後ろに追加
			text := inputs[1]
			tokens = pushBack(tokens, text)

		case "1": // 後ろを削除
			tokens = popBack(tokens)

		case "2": // 出力
			printTokens(tokens)
		}
	}

}

func pushBack(tokens []string, text string) []string {
	return append(tokens, text)
}

func popBack(tokens []string) []string {
	if len(tokens) == 0 {
		return tokens
	}
	return tokens[:len(tokens)-1]
}

func printTokens(tokens []string) {
	fmt.Println(strings.Join(tokens, " "))
}
