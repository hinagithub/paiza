// 問題
// 自然数 H, W, A, B が与えられます。縦に H 行、横に W 列で計 H * W 個の (A, B) の形式で文字列を出力してください。
// （AとBを、カンマと半角スペースで区切ってください。）
// ただし、横は | (半角スペース バーティカルライン 半角スペース) 区切りで、縦は = で区切って出力してください。
// また、縦の文字列間で = を出力する際は、その上の行と文字数がそろうように出力します。
// また、A と B は 9 けたになるように半角スペースを数値の前(右詰め)に埋めて出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var rowLen int

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")
	height, _ := strconv.Atoi(inputs[0])
	width, _ := strconv.Atoi(inputs[1])
	a, _ := strconv.Atoi(inputs[2])
	b, _ := strconv.Atoi(inputs[3])

	for i := 0; i < height; i++ {
		rowTexts := []string{}
		for j := 0; j < width; j++ {
			format := fmt.Sprintf("(%9d, %9d)", a, b)
			rowTexts = append(rowTexts, format)
		}
		row := strings.Join(rowTexts, " | ")
		fmt.Println(row)
		if i == 0 {
			rowLen = len(row)
		}
		if i != height-1 {
			fmt.Println(strings.Repeat("=", rowLen))
		}
	}

}
