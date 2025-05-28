// 問題
// 人事のあなたは、N 人の中から採用者を決定します。
// N 人のテストの点数はそれぞれ A_i (1 ≦ i ≦ N)です。
// テストの点数が K 点以上の人全員を採用したいのですが、
// 得点の高い方から M 人に辞退されてしまいました。
// あなたの採用することのできる最大の人数を答えてください。
// 採用できる人数が 0 人の場合もあることに気をつけてください。
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	inputs := strings.Fields(readline(scanner))
	applicantCount := atoi(inputs[0])
	border := atoi(inputs[1])
	withdrawnCount := atoi(inputs[2])

	var candidates []int
	for i := 0; i < applicantCount; i++ {
		score := atoi(readline(scanner))
		if score >= border {
			candidates = append(candidates, score)
		}
	}

	// 降順ソートし得点の高い方から M 人を除く
	// sort.IntSlice(candidates) []int はスライスをソート可能な型に変換 するもの
	sort.Sort(sort.Reverse(sort.IntSlice(candidates)))
	if len(candidates) < withdrawnCount {
		fmt.Println(0)
		return
	}
	fmt.Println(len(candidates) - withdrawnCount)
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
