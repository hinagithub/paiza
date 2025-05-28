// 問題
// あなたはボウリングをしています。フレームの 1 投目を投げ終わったあなたは、
// 次に狙うピンの番号と残っているピンの本数を知りたくなりました。
// ピンの情報が与えられるので、狙うべきピンの番号と残っているピンの本数を求めてください。

// 狙うピンの決め方は次の通りとします。
// - 残っているピンのうち、先頭 (P_1側) のピンを狙います。
// ただし、同じ列に複数ピンがある場合は、それらのうちピン番号が最も小さいピンを狙います。

// P_10 P_9 P_8 P_7
// P_6 P_5 P_4
// P_3 P_2
// P_1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// ピン番号のマトリクス（上から順）
	pinNumbers := [][]int{
		{10, 9, 8, 7},
		{6, 5, 4},
		{3, 2},
		{1},
	}

	pinStates := make([][]string, 4)
	for i := 0; i < 4; i++ {
		pinStates[i] = strings.Fields(readLine(scanner))
	}

	targetPin := 0
	remainCount := 0

	for i := 0; i < 4; i++ {
		for j := 0; j < len(pinStates[i]); j++ {
			if pinStates[i][j] == "1" {
				remainCount++
				pin := pinNumbers[i][j]
				if targetPin == 0 || pin < targetPin {
					targetPin = pin
				}
			}
		}
	}

	fmt.Println(targetPin)
	fmt.Println(remainCount)
}

func readLine(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "failed to read input")
		os.Exit(1)
	}
	return scanner.Text()
}
