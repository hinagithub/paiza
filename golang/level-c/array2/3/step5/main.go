// 問題
// 開店直後に店に入るために、番号 1 〜 N の N 人が開店前に店の前にブルーシートを合計 K 枚置きました。
// ブルーシートの先頭は A_1 , 最後尾は A_K です。
// しかし、店側は先頭から F 枚のブルーシートを撤去しました。

// 1 〜 N 番の人々は次のルールに従って店に並びます。
// ・自分のブルーシートが 1 枚以上残っているとき、
// 自分のブルーシートのうち、最も先頭に近いブルーシートの位置に並ぶ。
// ・自分のブルーシートが 1 枚も残っていないとき、並ぶことなく帰宅する。

// 全員が並び終わった後に待機列にいる人の番号を先頭から順に答えてください。

// 入力例
// 5 10 1
// 1
// 4
// 4
// 3
// 5
// 4
// 2
// 4
// 1
// 1
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

	// 1行目: N K F
	inputs := strings.Fields(readLine(scanner))
	// personCount := atoi(inputs[0]) // N 今回は使わない
	sheetCount := atoi(inputs[1])  // K
	removeCount := atoi(inputs[2]) // F

	seen := make(map[int]bool, sheetCount)
	queue := []int{}

	for i := 0; i < sheetCount; i++ {
		person := atoi(readLine(scanner))

		// F枚目までは撤去される
		if i < removeCount {
			continue
		}

		// まだ並んでいなければ並ばせる
		if !seen[person] {
			seen[person] = true
			queue = append(queue, person)
		}
	}

	// 出力
	for _, p := range queue {
		fmt.Println(p)
	}
}

func readLine(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "failed to scan input")
		os.Exit(1)
	}
	return scanner.Text()
}

func atoi(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to convert to int: %v\n", err)
		os.Exit(1)
	}
	return num
}
