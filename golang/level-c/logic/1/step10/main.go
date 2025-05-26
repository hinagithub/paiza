// 問題
// 0 または 1 の整数 A, B, C, D が与えられます。 以下の式を計算した結果を出力してください。

// NOT((NOTA AND NOTB) OR NOTC) XOR D

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 入力値（0 または 1）
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputs := strings.Fields(scanner.Text())
	A, _ := strconv.Atoi(inputs[0])
	B, _ := strconv.Atoi(inputs[1])
	C, _ := strconv.Atoi(inputs[2])
	D, _ := strconv.Atoi(inputs[3])

	// ビット反転（1ビットマスクを適用）
	notA := ^A & 1
	notB := ^B & 1
	notC := ^C & 1

	// 中間計算
	andResult := notA & notB
	orResult := andResult | notC
	notOrResult := ^orResult & 1

	// 最終結果
	result := notOrResult ^ D

	fmt.Println(result)
}
