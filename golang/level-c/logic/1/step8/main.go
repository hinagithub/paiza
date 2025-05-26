// 京子ちゃんは 二進数 1 けたの整数 A と B を使って、 A + B を計算したいと思っています。
// A と B を足した結果 (2 進表記) の下から 2 けた目の値を C 、下から 1 けた目の値を S とします。
// C と S を出力してください。
// 京子ちゃんは 二進数 1 けたの整数 A と B を使って、 A + B を計算したいと思っています。
// A と B を足した結果 (2 進表記) の下から 2 けた目の値を C 、下から 1 けた目の値を S とします。
// C と S を出力してください。
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

	sum := a + b

	stringSum := strconv.FormatInt(int64(sum), 2)

	if len(stringSum) < 2 {
		stringSum = "0" + stringSum
	}

	c := stringSum[len(stringSum)-2]
	s := stringSum[len(stringSum)-1]

	// 出力
	fmt.Println(string(c), string(s))
}
