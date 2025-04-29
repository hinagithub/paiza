// 問題
// 要素数 N の数列 A と要素数 Q の数列 B が与えられます。
// 1 ≦ i ≦ Q の各 i について、i 行目に A の B_i 番目の値を出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	scanner.Scan() // 1行目: N（使わない）
	scanner.Scan()
	texts := strings.Fields(scanner.Text())

	scanner.Scan() // 3行目: Q（使わない）
	scanner.Scan()
	targets := strings.Fields(scanner.Text())

	for _, target := range targets {
		// インデックスを取得
		index, err := strconv.Atoi(target)
		if err != nil || index < 1 || index > len(texts) {
			fmt.Println("invalid index:", target)
			continue
		}
		fmt.Println(texts[index-1])
	}
}
