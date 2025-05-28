// 問題
// データ構造の queue と同様の働きをするロボットがあります。
// ロボットは指示に応じて配列 A に対して 2 種類の仕事を行います、仕事の内容は以下の通りです。

// ・in n と指示された場合、A の末尾に n を追加してください。
// ・out と指示された場合、A の先頭の要素を削除してください。ただし、A が既に空の場合、何も行わないでください。

// ロボットに与えられる指示の回数 N と、各指示の内容 S_i が与えられるので、
// ロボットが全ての処理を順に行った後の A の各要素を出力してください。
// なお、初め配列 A は空であるものとします。

// 入力例
// 10
// out
// in 33
// out
// out
// out
// out
// in -76
// out
// out
// in -53

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	const (
		In  = "in"
		Out = "out"
	)
	scanner := bufio.NewScanner(os.Stdin)
	orderCount := atoi(readline(scanner))
	arr := []int{}

	for i := 0; i < orderCount; i++ {
		inputs := strings.Fields(readline(scanner))
		order := inputs[0]
		if order == In {
			val := atoi(inputs[1])
			arr = in(arr, val)
		} else {
			arr = out(arr)
		}
	}

	for _, v := range arr {
		fmt.Println(v)
	}

}

func in(arr []int, num int) []int {
	arr = append(arr, num)
	return arr
}

func out(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	return arr[1:]
}

func readline(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "failed to scan input")
		os.Exit(1)
	}
	return scanner.Text()
}

func atoi(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to atoi: %v\n", err)
		os.Exit(1)
	}
	return num
}
