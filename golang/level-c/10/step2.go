// 自然数 N, A, B が与えられます。
// (A, B) という形式の文字列を N 回、カンマと半角スペース区切りで出力してください。

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
	inputs := strings.Split(sc.Text(), " ")
	len, _ := strconv.Atoi(inputs[0])
	a, _ := strconv.Atoi(inputs[1])
	b, _ := strconv.Atoi(inputs[2])
	results := make([]string, len)
	for i := 0; i < len; i++ {
		results[i] = fmt.Sprintf("(%d, %d)", a, b)
	}
	fmt.Println(strings.Join(results, ", "))
}
