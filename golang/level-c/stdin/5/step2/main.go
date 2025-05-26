//　問題
// 整数 a, b が半角スペース区切りで与えられるので、改行区切りにして 2 行で出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")
	fmt.Println(inputs[0])
	fmt.Println(inputs[1])
}
