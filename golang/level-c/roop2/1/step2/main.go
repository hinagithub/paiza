// 問題
// 複数の文字列が入力されます。文字列の数はわかりません。
// EOF が入力されるまで、受け取った文字列を改行区切りで出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texts := strings.Fields(scanner.Text())
	for _, text := range texts {
		fmt.Println(text)
		if text == "EOL" {
			return
		}
	}
}
