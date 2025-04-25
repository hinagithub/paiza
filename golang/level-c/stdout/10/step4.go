// 問題
// 自然数 H, W, A, B が与えられます。縦に H 行、横に W 行で計 H * W 個の (A, B) という形式の文字列を出力してください。
// ただし、横は | (半角スペース バーティカルライン 半角スペース) 区切りで、縦は = で区切って出力してください。
// また、縦の文字列間で = を出力する際は、その上の行と文字数が等しくなるように出力します。

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
	height, _ := strconv.Atoi(inputs[0])
	width, _ := strconv.Atoi(inputs[1])
	a, _ := strconv.Atoi(inputs[2])
	b, _ := strconv.Atoi(inputs[3])

	max := height - 1
	var lineLen = 0
	for i := 0; i < height; i++ {
		line := []string{}
		for j := 0; j < width; j++ {
			text := fmt.Sprintf("(%d, %d)", a, b)
			line = append(line, text)
		}
		result := strings.Join(line, " | ")
		fmt.Println(result)
		lineLen = len(result)

		if i != max {
			fmt.Println(strings.Repeat("=", lineLen))
		}
	}
}
