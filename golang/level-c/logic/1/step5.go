// 問題
// 0 または 1 の整数 A と B が与えられます。 A NAND B の結果を出力してください。

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
	scanner.Scan()

	inputs := strings.Fields(scanner.Text())
	a, _ := strconv.Atoi(inputs[0])
	b, _ := strconv.Atoi(inputs[1])

	if a == 1 && b == 1 {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}
}
