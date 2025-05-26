// 問題
// 1 行目で整数 N が与えられます。
// 2 行目で、N 個の整数 a_1, ... , a_N が半角スペース区切りで与えられます。
// a_1, ... , a_N を改行区切りで出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan() // 1行目は読み飛ばす
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")
	for _, input := range inputs {
		fmt.Println(input)
	}
}
