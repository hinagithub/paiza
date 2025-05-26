// 問題
// K さんは paiza のアカウントを作成することにしました。
// そのためには、パスワードの設定が必要なことがわかりました。
// そこで K さんは忘れないように、次のようなルールにのっとって
// N 文字のパスワードを設定することにしました。

// ・ ルール
// K さんは N 文字のうち、 Q 文字だけ覚えておく文字を決めておく。
// 具体的には n_i 文字目を c_i とだけ決めて、残りの全ての文字を C にする。

// K さんの設定したパスワードを当ててください。

// 例
// ・ N = 5 , Q = 1 , n_1 = 2 , c_1 = 'T' , C = 'K' のとき
// パスワードは 5 文字であり、 2 文字目が 'T' ,
// それ以外の文字を 'K' としたものである "KTKKK" が K さんのパスワードとなる。

// 入力例
// 5
// 1
// 2 T
// K

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	length := atoi(readline(scanner))
	altLength := atoi(readline(scanner))
	alts := make(map[int]string)

	for i := 0; i < altLength; i++ {
		inputs := strings.Fields(readline(scanner))
		index := atoi(inputs[0]) - 1
		text := inputs[1]
		alts[index] = text
	}

	defaultChar := readline(scanner)
	password := make(map[int]string)
	for i := 0; i < length; i++ {
		if ch, ok := alts[i]; ok {
			password[i] = ch
		} else {
			password[i] = defaultChar
		}
	}
	fmt.Println(joinValues(password))

}

func readline(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Println("scan err")
		os.Exit(1)
	}
	return scanner.Text()
}

func atoi(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "atoi err", err)
		os.Exit(1)
	}
	return num
}

func joinValues(keyValues map[int]string) string {
	// キーを昇順に並べる
	keys := make([]int, 0, len(keyValues))
	for k := range keyValues {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// 配列にして返却
	result := ""
	for _, key := range keys {
		result = result + keyValues[key]
	}
	return result
}
