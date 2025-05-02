// 問題
// 0 または 1 の整数 A と B が与えられます。 A XOR B の結果を出力してください。

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

	fmt.Println(a ^ b)
}
