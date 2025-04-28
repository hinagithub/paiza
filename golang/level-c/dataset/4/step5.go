// 問題
// N 個の要素からなる数列 A, B が与えられます。A または B に含まれる値をすべて列挙し、重複を取り除いて昇順に出力してください。

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
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan() // 1行目の文字数Nは使わないので読み飛ばす

	// AとBの配列を取得
	sc.Scan()
	aInputs := strings.Fields(sc.Text())
	sc.Scan()
	bInputs := strings.Fields(sc.Text())

	// まとめて数値に変換
	numbers := append(AtoiAll(aInputs), AtoiAll(bInputs)...)

	// 重複排除
	unique := make(map[int]bool)
	for _, num := range numbers {
		unique[num] = true
	}

	// mapからスライスに変換してソート
	results := make([]int, 0, len(unique))
	for num := range unique {
		results = append(results, num)
	}
	sort.Ints(results)

	// 文字列を出力
	resultTexts := ItoaAll(results)
	fmt.Println(strings.Join(resultTexts, " "))

}

// 配列内の文字列を数値に置き換える
func AtoiAll(texts []string) []int {
	numbers := []int{}
	for _, text := range texts {
		number, _ := strconv.Atoi(text)
		numbers = append(numbers, number)
	}
	return numbers
}

// 配列内の数値を文字列に置き換える
func ItoaAll(numbers []int) []string {
	texts := []string{}
	for _, num := range numbers {
		text := strconv.Itoa(num)
		texts = append(texts, text)
	}
	return texts
}

// 配列内で一致する数値があれば1、なければ0を返却する
func FindExists(numbers []int, targetNum int) bool {
	for _, num := range numbers {
		if num == targetNum {
			return true
		}
	}
	return false
}
