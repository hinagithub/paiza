// 整数 A, B, C が与えられます。
// 式 A × B ≦ C が成立している場合はYESを、そうではない場合はNOを出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Text()
	inputs := strings.Split(text, " ")
	a, _ := strconv.Atoi(inputs[0])
	b, _ := strconv.Atoi(inputs[1])
	c, _ := strconv.Atoi(inputs[2])

	if a*b <= c {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
