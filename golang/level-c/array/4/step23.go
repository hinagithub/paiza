// 問題
// 5 つの文字列 s_1, s_2, ..., s_5 が半角スペース区切りで与えられます。
// 5 つの文字列を辞書順に並べ替え、改行区切りで出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	strs := strings.Fields(scanner.Text())
	sort.Strings(strs)
	for _, str := range strs {
		fmt.Println(str)
	}
}
